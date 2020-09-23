package internal

import (
	"context"
	"path"
	"path/filepath"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/packethost/pkg/log"
	"github.com/pkg/errors"
	pb "github.com/tinkerbell/tink/protos/workflow"
)

const (
	errCreateContainer = "failed to create container"
	errFailedToWait    = "failed to wait for completion of action"
	errFailedToRunCmd  = "failed to run on-timeout command"

	infoWaitFinished = "wait finished for failed or timeout container"
)

func (w *Worker) createContainer(ctx context.Context, cmd []string, wfID string, action *pb.WorkflowAction) (string, error) {
	registry := w.registry
	config := &container.Config{
		Image:        path.Join(registry, action.GetImage()),
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          cmd,
		Tty:          true,
		Env:          action.GetEnvironment(),
	}

	wfDir := filepath.Join(dataDir, wfID)
	hostConfig := &container.HostConfig{
		Privileged: true,
		Binds:      []string{wfDir + ":/workflow"},
	}
	hostConfig.Binds = append(hostConfig.Binds, action.GetVolumes()...)
	w.logger.With("command", cmd).Info("creating container")
	resp, err := w.registryClient.ContainerCreate(ctx, config, hostConfig, nil, action.GetName())
	if err != nil {
		return "", errors.Wrap(err, "DOCKER CREATE")
	}
	return resp.ID, nil
}

func startContainer(ctx context.Context, l log.Logger, cli *client.Client, id string) error {
	l.With("containerID", id).Debug("starting container")
	return errors.Wrap(cli.ContainerStart(ctx, id, types.ContainerStartOptions{}), "DOCKER START")
}

func waitContainer(ctx context.Context, cli *client.Client, id string) (pb.ActionState, error) {
	// Inspect whether the container is in running state
	if _, err := cli.ContainerInspect(ctx, id); err != nil {
		return pb.ActionState_ACTION_FAILED, nil
	}

	// send API call to wait for the container completion
	wait, errC := cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)

	select {
	case status := <-wait:
		if status.StatusCode == 0 {
			return pb.ActionState_ACTION_SUCCESS, nil
		}
		return pb.ActionState_ACTION_FAILED, nil
	case err := <-errC:
		return pb.ActionState_ACTION_FAILED, err
	case <-ctx.Done():
		return pb.ActionState_ACTION_TIMEOUT, ctx.Err()
	}
}

func waitFailedContainer(ctx context.Context, l log.Logger, cli *client.Client, id string, failedActionStatus chan pb.ActionState) {
	// send API call to wait for the container completion
	wait, errC := cli.ContainerWait(ctx, id, container.WaitConditionNotRunning)

	select {
	case status := <-wait:
		if status.StatusCode == 0 {
			failedActionStatus <- pb.ActionState_ACTION_SUCCESS
		}
		failedActionStatus <- pb.ActionState_ACTION_FAILED
	case err := <-errC:
		l.Error(err)
		failedActionStatus <- pb.ActionState_ACTION_FAILED
	case <-ctx.Done():
		l.Error(ctx.Err())
		failedActionStatus <- pb.ActionState_ACTION_TIMEOUT
	}
}

func removeContainer(ctx context.Context, l log.Logger, cli *client.Client, id string) error {
	// create options for removing container
	opts := types.ContainerRemoveOptions{
		Force:         true,
		RemoveLinks:   false,
		RemoveVolumes: true,
	}
	l.With("containerID", id).Info("removing container")

	// send API call to remove the container
	return cli.ContainerRemove(ctx, id, opts)
}
