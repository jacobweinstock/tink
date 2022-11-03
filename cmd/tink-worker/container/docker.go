package container

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/tinkerbell/tink/protos/workflow"
)

type Docker struct {
	Conn         conn
	RegistryAuth map[string]string
}

type conn interface {
	client.ContainerAPIClient
	client.ImageAPIClient
}

// configOpt allows modifying the container config defaults.
type configOpt func(*container.Config)

// hostOpt allows modifying the container host config defaults.
type hostOpt func(*container.HostConfig)

func (d *Docker) Prepare(ctx context.Context, action *workflow.WorkflowAction) (id string, err error) {
	// 1. Pull the image
	if pullErr := d.pullImage(ctx, action.GetImage(), types.ImagePullOptions{RegistryAuth: getRegistryAuth(d.RegistryAuth, action.GetImage())}); pullErr != nil {
		return "", pullErr
	}
	// 2. create container
	containerName := fmt.Sprintf("%v-%v", strings.ReplaceAll(action.Name, " ", "-"), time.Now().UnixNano())
	id, err = d.createContainer(ctx, containerName, toDockerConf(ctx, action), toDockerHostConfig(ctx, action))
	if err != nil {
		return "", err
	}

	return id, nil
}

func (d *Docker) Run(ctx context.Context, id string, logs chan []byte) error {
	if err := d.Conn.ContainerStart(ctx, id, types.ContainerStartOptions{}); err != nil {
		return err
	}

	// stream logs
	go func() {
		_ = d.streamLogs(ctx, id, types.ContainerLogsOptions{ShowStdout: true, ShowStderr: true, Follow: true}, logs)
	}()

	var detail types.ContainerJSON
	var complete bool
	for !complete {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			var ok bool
			var err error
			ok, detail, err = d.containerExecComplete(ctx, id)
			if err != nil {
				continue
			}
			if ok {
				complete = true
			}
		}
	}

	if detail.ContainerJSONBase == nil {
		return errors.New("container details was nil, cannot tell success or failure status without these details")
	}
	// container execution completed successfully
	if detail.State.ExitCode == 0 {
		return nil
	}

	return fmt.Errorf("msg: container execution was unsuccessful; exitCode: %v; details: %v", detail.State.ExitCode, detail.State.Error)
}

func (d *Docker) Destroy(ctx context.Context, id string, timeout time.Duration) error {
	select {
	case <-ctx.Done(): // a canceled context will cause ContainerStop and ContainerRemove methods to not run.
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
		defer cancel()
	default:
	}
	_ = d.Conn.ContainerStop(ctx, id, nil) // stopping before removing allows any final logging messages to be captured // TODO: maybe log the error? this would require the Docker struct to have a logger field.

	return d.Conn.ContainerRemove(ctx, id, types.ContainerRemoveOptions{Force: true, RemoveVolumes: true})
}

// pullImage is what you would expect from a `docker pull` cli command
// pulls an image from a remote registry.
func (d *Docker) pullImage(ctx context.Context, image string, pullOpts types.ImagePullOptions) error {
	out, err := d.Conn.ImagePull(ctx, image, pullOpts)
	if err != nil {
		return errors.Wrapf(err, "error pulling image: %v", image)
	}
	defer out.Close()
	fd := json.NewDecoder(out)
	var imagePullStatus struct {
		Error string `json:"error"`
	}
	for {
		if err := fd.Decode(&imagePullStatus); err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return errors.Wrapf(err, "error pulling image: %v", image)
		}
		if imagePullStatus.Error != "" {
			return errors.Wrapf(errors.New(imagePullStatus.Error), "error pulling image: %v", image)
		}
	}

	return nil
}

// createContainer creates a container that is not started.
func (d *Docker) createContainer(ctx context.Context, containerName string, containerConfig *container.Config, hostConfig *container.HostConfig) (id string, err error) {
	resp, err := d.Conn.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, containerName)
	if err != nil {
		return "", err
	}

	return resp.ID, nil
}

// streamLogs streams the logs of a container to a []byte channel. Is blocking until the .Scan() function returns.
func (d *Docker) streamLogs(_ context.Context, containerID string, options types.ContainerLogsOptions, logs chan []byte) error {
	reader, err := d.Conn.ContainerLogs(context.Background(), containerID, options)
	if err != nil {
		return err
	}
	defer reader.Close()

	buf := bufio.NewScanner(reader)
	for buf.Scan() {
		logs <- buf.Bytes()
	}

	return buf.Err()
}

// containerExecComplete checks if a container run has completed or not. completed is defined as having an "exited" or "dead" status.
// see types.ContainerJSON.State.Status for all status options.
func (d *Docker) containerExecComplete(ctx context.Context, containerID string) (complete bool, details types.ContainerJSON, err error) {
	detail, err := d.Conn.ContainerInspect(ctx, containerID)
	if err != nil {
		return false, types.ContainerJSON{}, errors.Wrap(err, "unable to inspect container")
	}

	if detail.State.Status == "exited" || detail.State.Status == "dead" {
		return true, detail, nil
	}

	return false, types.ContainerJSON{}, nil
}

func getRegistryAuth(regAuth map[string]string, imageName string) string {
	for reg, auth := range regAuth {
		if strings.HasPrefix(imageName, reg) {
			return auth
		}
	}

	return ""
}

// toDockerConf takes a workflowAction and translates it to a docker container config.
func toDockerConf(_ context.Context, workflowAction *workflow.WorkflowAction, opts ...configOpt) *container.Config {
	defaultConfig := &container.Config{
		AttachStdout: true,
		AttachStderr: true,
		Tty:          true,
		Env:          workflowAction.Environment,
		Cmd:          workflowAction.Command,
		Image:        workflowAction.Image,
	}
	for _, opt := range opts {
		opt(defaultConfig)
	}

	return defaultConfig
}

// toDockerHostConfig converts a tink action spec to a container host config spec.
func toDockerHostConfig(_ context.Context, workflowAction *workflow.WorkflowAction, opts ...hostOpt) *container.HostConfig {
	defaultConfig := &container.HostConfig{
		Binds:      workflowAction.Volumes,
		PidMode:    container.PidMode(workflowAction.Pid),
		Privileged: true,
	}
	for _, opt := range opts {
		opt(defaultConfig)
	}

	return defaultConfig
}
