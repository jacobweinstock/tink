package worker

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-logr/logr"
	"github.com/tinkerbell/tink/protos/workflow"
)

/*
func actionController(ctx context.Context, logger logr.Logger, queue <-chan *workflow.WorkflowAction) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("context cancelled")
			return
		case action := <-queue:
			// TODO: implement
			logger.Info("received action", "action", fmt.Sprintf("%+v", action))
			time.Sleep(time.Second)
		}
	}
}

func statusReportController(ctx context.Context, logger logr.Logger, queue <-chan *workflow.WorkflowActionStatus) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("context cancelled")
			return
		case status := <-queue:
			// TODO: implement
			logger.Info("received action", "action", fmt.Sprintf("%+v", status))
			time.Sleep(time.Second)
		}
	}
}
*/

// Run executes an action until successful or the context is cancelled.
func Run(ctx context.Context, logger logr.Logger, cm ContainerRunner, action *workflow.WorkflowAction) {
	for {
		select {
		case <-ctx.Done():
			logger.Info("context cancelled")
			return
		default:
			if err := executeAction(ctx, logger, cm, action); err != nil {
				logger.Error(err, "action execution failed")
				time.Sleep(time.Second * 2) // wait before retrying // TODO: make configurable
				continue
			}
			return
		}
	}
}

// ContainerRunner defines the methods needed to run a workflow task action.
type ContainerRunner interface {
	// Prepare should pull images, create (not run) any containers/pods, setup the environment, mounts, namespaces, etc
	Prepare(ctx context.Context, action *workflow.WorkflowAction) (id string, err error)
	// Run should execution the action and wait for completion
	Run(ctx context.Context, id string, logs chan []byte) error
	// Destroy should handle removing all things created/setup in Prepare
	Destroy(ctx context.Context, id string, timeout time.Duration) error
}

func executeAction(ctx context.Context, logger logr.Logger, cm ContainerRunner, action *workflow.WorkflowAction) error {
	// 1. Pull image(prepare phase)
	// 2. Create container(prepare phase)
	// 3. defer container delete(destroy phase)
	// 4. Stream logs to stdout(this function)
	// 5. Start container(run phase)
	// 6. Wait for container to finish(run phase)
	id, err := cm.Prepare(ctx, action)
	if err != nil {
		return err
	}
	defer func() {
		// TODO: make timeout configurable
		if err := cm.Destroy(ctx, id, time.Second*5); err != nil {
			logger.Error(err, "failed to destroy container")
		}
	}()

	logs := make(chan []byte)
	go streamLogs(ctx, logger, id, logs, action)

	return cm.Run(ctx, id, logs)
}

func streamLogs(ctx context.Context, logger logr.Logger, id string, logs chan []byte, action *workflow.WorkflowAction) {
	for {
		select {
		case <-ctx.Done():
		case log := <-logs:
			kvs := []interface{}{
				"action", action.String(),
				"log", string(log),
				"id", id,
			}
			if b, err := json.Marshal(action); err == nil {
				act := make(map[string]interface{})
				if err := json.Unmarshal(b, &act); err != nil {
					logger.Error(err, "failed to unmarshal action")
				}
				kvs[1] = act
			}
			m := make(map[string]interface{})
			if err := json.Unmarshal(log, &m); err == nil {
				kvs[3] = m
			}
			logger.Info("container logs", kvs...)
		}
	}
}
