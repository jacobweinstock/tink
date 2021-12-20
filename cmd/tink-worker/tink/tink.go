package tink

import (
	"context"

	"github.com/pkg/errors"
	pb "github.com/tinkerbell/tink/protos/workflow"
)

type WorkflowData struct{}

// WorkflowContexts gets workflow context for worker id on success else returns error.
func (w WorkflowData) WorkflowContexts(ctx context.Context, client pb.WorkflowServiceClient, workerID string) (pb.WorkflowService_GetWorkflowContextsClient, error) {
	if workerID == "" {
		return nil, errors.New("empty string is not a valid worker id")
	}

	if client == nil {
		return nil, errors.New("nil WorkflowServiceClient is not a valid interface")
	}

	response, err := client.GetWorkflowContexts(ctx, &pb.WorkflowContextRequest{WorkerId: workerID})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get workflow contexts")
	}
	return response, nil
}

// WorkflowActions gets workflow action list for workflow id on success else returns error.
func (w WorkflowData) WorkflowActions(ctx context.Context, client pb.WorkflowServiceClient, workflowID string) (*pb.WorkflowActionList, error) {
	if workflowID == "" {
		return nil, errors.New("empty string is not a valid workflow id")
	}

	if client == nil {
		return nil, errors.New("nil WorkflowServiceClient is not a valid interface")
	}

	actions, err := client.GetWorkflowActions(ctx, &pb.WorkflowActionsRequest{WorkflowId: workflowID})
	if err != nil {
		return nil, errors.Wrap(err, "failed to get workflow actions")
	}
	return actions, nil
}

// ReportWorkflowActionStatus reports action status on success else returns error.
func (w WorkflowData) ReportWorkflowActionStatus(ctx context.Context, client pb.WorkflowServiceClient, actionStatus *pb.WorkflowActionStatus) error {
	if client == nil {
		return errors.New("nil WorkflowServiceClient is not a valid interface")
	}

	if actionStatus == nil {
		return errors.New("nil WorkflowActionStatus is not a valid action status")
	}

	_, err := client.ReportActionStatus(ctx, actionStatus)
	if err != nil {
		return errors.Wrap(err, "failed to report action status")
	}
	return err
}
