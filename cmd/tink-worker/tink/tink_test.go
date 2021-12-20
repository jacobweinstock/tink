package tink

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/pkg/errors"
	pb "github.com/tinkerbell/tink/protos/workflow"
	"google.golang.org/grpc"
)

type mockCli struct {
	pb.WorkflowServiceClient
	response     pb.WorkflowService_GetWorkflowContextsClient
	actionstatus pb.WorkflowActionStatus
	actionlist   *pb.WorkflowActionList
}

func (mcli *mockCli) GetWorkflowContexts(ctx context.Context, _ *pb.WorkflowContextRequest, _ ...grpc.CallOption) (pb.WorkflowService_GetWorkflowContextsClient, error) {
	if ctx == nil {
		return nil, errors.New("nil context is not a valid interface")
	}

	return mcli.response, nil
}

func (mcli *mockCli) GetWorkflowActions(ctx context.Context, _ *pb.WorkflowActionsRequest, _ ...grpc.CallOption) (*pb.WorkflowActionList, error) {
	if ctx == nil {
		return nil, errors.New("nil context is not a valid interface")
	}
	return mcli.actionlist, nil
}

func (mcli *mockCli) ReportActionStatus(ctx context.Context, _ *pb.WorkflowActionStatus, _ ...grpc.CallOption) (*pb.Empty, error) {
	if ctx == nil {
		return nil, errors.New("nil context is not a valid interface")
	}
	return nil, nil
}

func TestWorkerGetWorkflowActions(t *testing.T) {
	mock := &mockCli{}
	type args struct {
		ctx        context.Context
		cli        pb.WorkflowServiceClient
		workflowID string
	}
	tests := []struct {
		name    string
		data    args
		w       WorkflowData
		wantErr error
	}{
		{
			name: "get_workflow_actions_with_workflowID",
			data: args{
				ctx:        context.Background(),
				cli:        mock,
				workflowID: "3431423",
			},
			wantErr: nil,
		},
		{
			name: "get_workflow_actions_with_no_workflowID",
			data: args{
				ctx:        context.Background(),
				cli:        mock,
				workflowID: "",
			},
			wantErr: errors.New("empty string is not a valid workflow id"),
		},
		{
			name: "get_workflow_actions_with_no_service_client",
			data: args{
				ctx:        context.Background(),
				cli:        nil,
				workflowID: "3431423",
			},
			wantErr: errors.New("nil WorkflowServiceClient is not a valid interface"),
		},
		{
			name: "get_workflow_actions_with_service_client",
			data: args{
				ctx:        context.Background(),
				cli:        mock,
				workflowID: "3431423",
			},
			wantErr: nil,
		},
		{
			name: "get_workflow_actions_with_no_context",
			data: args{
				ctx:        nil,
				cli:        mock,
				workflowID: "3431423",
			},
			wantErr: errors.New("failed to get workflow actions: nil context is not a valid interface"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.w.WorkflowActions(tt.data.ctx, tt.data.cli, tt.data.workflowID)
			if err != nil {
				diff := cmp.Diff(tt.wantErr.Error(), err.Error())
				if diff != "" {
					t.Fatal(diff)
				}
			}
		})
	}
}

func TestWorkerReportActionStatus(t *testing.T) {
	mock := &mockCli{}
	type args struct {
		ctx          context.Context
		cli          pb.WorkflowServiceClient
		actionStatus *pb.WorkflowActionStatus
	}
	tests := []struct {
		name    string
		data    args
		w       WorkflowData
		wantErr error
	}{
		{
			name: "get_worker_report_action_status_with_no_service_client",
			data: args{
				ctx:          context.Background(),
				cli:          nil,
				actionStatus: &mock.actionstatus,
			},
			wantErr: errors.New("nil WorkflowServiceClient is not a valid interface"),
		},
		{
			name: "get_worker_report_action_status_with_service_client",
			data: args{
				ctx:          context.Background(),
				cli:          mock,
				actionStatus: &mock.actionstatus,
			},
			wantErr: nil,
		},
		{
			name: "get_worker_report_action_status_with_no_action_status",
			data: args{
				ctx:          context.Background(),
				cli:          mock,
				actionStatus: nil,
			},
			wantErr: errors.New("nil WorkflowActionStatus is not a valid action status"),
		},
		{
			name: "get_worker_report_action_status_with_action_status",
			data: args{
				ctx:          context.Background(),
				cli:          mock,
				actionStatus: &mock.actionstatus,
			},
			wantErr: nil,
		},
		{
			name: "get_worker_report_action_status_with_no_context",
			data: args{
				ctx:          nil,
				cli:          mock,
				actionStatus: &mock.actionstatus,
			},
			wantErr: errors.New("failed to report action status: nil context is not a valid interface"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.w.ReportWorkflowActionStatus(tt.data.ctx, tt.data.cli, tt.data.actionStatus)
			if err != nil {
				diff := cmp.Diff(tt.wantErr.Error(), err.Error())
				if diff != "" {
					t.Fatal(diff)
				}
			}
		})
	}
}

func TestWorkerGetWorkflowContexts(t *testing.T) {
	mock := &mockCli{}
	type args struct {
		ctx      context.Context
		cli      pb.WorkflowServiceClient
		workerID string
	}
	tests := []struct {
		name    string
		data    args
		w       WorkflowData
		wantErr error
	}{
		{
			name: "get_worker_flow_context_with_no_service_client",
			data: args{
				ctx:      context.Background(),
				cli:      nil,
				workerID: "3431423",
			},
			wantErr: errors.New("nil WorkflowServiceClient is not a valid interface"),
		},
		{
			name: "get_worker_flow_context_with_service_client",
			data: args{
				ctx:      context.Background(),
				cli:      mock,
				workerID: "3431423",
			},
			wantErr: nil,
		},
		{
			name: "get_worker_flow_context_with_workerID",
			data: args{
				ctx:      context.Background(),
				cli:      mock,
				workerID: "3431423",
			},
			wantErr: nil,
		},
		{
			name: "get_worker_flow_context_with_no_workerID",
			data: args{
				ctx:      context.Background(),
				cli:      mock,
				workerID: "",
			},
			wantErr: errors.New("empty string is not a valid worker id"),
		},
		{
			name: "get_worker_flow_context_with_no_context",
			data: args{
				ctx:      nil,
				cli:      mock,
				workerID: "3431423",
			},
			wantErr: errors.New("failed to get workflow contexts: nil context is not a valid interface"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.w.WorkflowContexts(tt.data.ctx, tt.data.cli, tt.data.workerID)
			if err != nil {
				diff := cmp.Diff(tt.wantErr.Error(), err.Error())
				if diff != "" {
					t.Fatal(diff)
				}
			}
		})
	}
}
