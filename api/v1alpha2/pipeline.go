package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

const (
	// PipelineStatePreparing indicates the Pipeline is preparing to run.
	PipelineStatePreparing PipelineState = "Preparing"

	// PipelineStatePending indicates the Pipeline is awaiting dispatch to the agent.
	PipelineStatePending PipelineState = "Pending"

	// PipelineStateScheduled indicates the Pipeline has been dispatched to the agent but the agent
	// is yet to report the Pipeline has started.
	PipelineStateScheduled PipelineState = "Scheduled"

	// PipelineStateRunning indicates the Pipeline has begun executing.
	PipelineStateRunning PipelineState = "Running"

	// PipelineStateCancelling indicates the agent has been instructed to cancel the Pipeline, but
	// the cancellation is yet to be completed.
	PipelineStateCancelling PipelineState = "Cancelling"

	// PipelineStateSucceeded indicates all Actions have successfully completed.
	PipelineStateSucceeded PipelineState = "Succeeded"

	// PipelineStateFailed indicates an Action failed. This includes timeouts and agent side failures.
	PipelineStateFailed PipelineState = "Failed"

	// PipelineStateCanceled indicates the Pipeline has been canceled.
	PipelineStateCanceled PipelineState = "Canceled"

	// Action State constants

	// ActionStatePending indicates an Action is awaiting execution.
	ActionStatePending ActionState = "Pending"

	// ActionStateRunning indicates an Action has begun execution.
	ActionStateRunning ActionState = "Running"

	// ActionStateSucceeded indicates an Action completed execution successfully.
	ActionStateSucceeded ActionState = "Succeeded"

	// ActionStatFailed indicates an Action failed to execute. Users may inspect the associated
	// Workflow resource to gain deeper insights into why the action failed.
	ActionStateFailed ActionState = "Failed"

	// Pipeline Condition constants

	NetbootJobFailed        PipelineConditionType = "NetbootJobFailed"
	NetbootJobComplete      PipelineConditionType = "NetbootJobComplete"
	NetbootJobRunning       PipelineConditionType = "NetbootJobRunning"
	NetbootJobSetupFailed   PipelineConditionType = "NetbootJobSetupFailed"
	NetbootJobSetupComplete PipelineConditionType = "NetbootJobSetupComplete"
	ToggleAllowNetbootTrue  PipelineConditionType = "AllowNetbootTrue"
	ToggleAllowNetbootFalse PipelineConditionType = "AllowNetbootFalse"
	WorkflowRenderedSuccess PipelineConditionType = "WorkflowRenderedSuccess"

	// Workflow Rendering constants

	WorkflowRenderingSucceeded WorkflowRenderingState = "Succeeded"
	WorkflowRenderingFailed    WorkflowRenderingState = "Failed"
	WorkflowRenderingUnknown   WorkflowRenderingState = "Unknown"

	// BootMode constants

	BootModeNetboot BootMode = "netboot"
	BootModeISO     BootMode = "isoboot"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:categories=tinkerbell,shortName=pl
// +kubebuilder:printcolumn:name="Hardware",type="string",JSONPath=".status.currentHardware",description="Hardware that is currently that corresponds to the current Agent and Workflow"
// +kubebuilder:printcolumn:name="Workflow",type="string",JSONPath=".status.CurrentWorkflow",description="Workflow that is currently be run by the associated Agent"
// +kubebuilder:printcolumn:name="AgentID",type="string",JSONPath=".status.CurrentAgent",description="ID of the Agent that is running the current Workflow"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state",description="Overall state of the Pipeline such as Pending,Running etc"
// +kubebuilder:printcolumn:name="Action",type="string",JSONPath=".status.currentAction",description="Action that currently being executed"

// Pipeline describes a set of actions to be run on a specific Hardware. Workflows execute
// once and should be considered ephemeral.
type Pipeline struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineSpec   `json:"spec,omitempty"`
	Status PipelineStatus `json:"status,omitempty"`
}

type PipelineWorkflow struct {
	// AgentID is a unique identifier of the Agent that will execute the Workflow.
	// +required
	AgentID string `json:"agentId,omitempty"`

	// BootOptions are options that control the booting of Hardware. A HardwareRef must be provided to use BootOptions.
	// +optional
	BootOptions *BootOptions `json:"bootOptions,omitempty"`

	// WorkflowRef is a reference to a Workflow resource used to render Workflow actions.
	// +required
	WorkflowRef corev1.LocalObjectReference `json:"workflowRef,omitempty"`

	// HardwareRef is a reference to a Hardware resource this workflow will execute on.
	// +optional
	HardwareRef corev1.LocalObjectReference `json:"hardwareRef,omitempty"`

	// TemplateParams are a list of key-value pairs that are injected into Workflows at render
	// time. TemplateParams are exposed to Workflow using a top level .Params key.
	//
	// For example, TemplateParams = {"foo": "bar"}, the foo key can be accessed via .Params.foo.
	// +optional
	TemplateParams map[string]string `json:"templateParams,omitempty"`

	// TimeoutSeconds defines the time the workflow has to complete. The timer begins when the first
	// action is requested. When set to 0, no timeout is applied.
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	TimeoutSeconds int64 `json:"timeout,omitempty"`
}

type PipelineSpec struct {
	// BootOptions are options that control the booting of Hardware for all Workflows. A HardwareRef in each Workflow must be provided to use BootOptions.
	// +optional
	BootOptions *BootOptions `json:"bootOptions,omitempty"`

	// TemplateParams are a list of key-value pairs that are injected into all Workflows at render
	// time. TemplateParams are exposed to Workflows using a top level .Params key.
	//
	// For example, TemplateParams = {"foo": "bar"}, the foo key can be accessed in a Workflow via .Params.foo.
	// +optional
	TemplateParams map[string]string `json:"templateParams,omitempty"`

	// TimeoutSeconds defines the time the Pipeline has to complete. The timer begins when the first
	// Action of the first Workflow is requested. When set to 0, no timeout is applied.
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	TimeoutSeconds int64 `json:"timeout,omitempty"`

	// Workflows are a list of workflows to be executed
	Workflows []PipelineWorkflow `json:"workflows,omitempty"`
}

type PipelineStatus struct {
	// Workflows is a list of action states.
	Workflows []ActionStatus `json:"actions"`

	// StartedAt is the time the first action was requested. Nil indicates the Workflow has not
	// started.
	// +optional
	StartedAt *metav1.Time `json:"startedAt,omitempty"`

	// LastTransition is the observed time when State transitioned last.
	LastTransition metav1.Time `json:"lastTransitioned,omitempty"`

	// State describes the current state of the Pipeline.
	State PipelineState `json:"state,omitempty"`

	// CurrentWorkflow is the workflow that is currently in the running state.
	CurrentWorkflow string `json:"currentWorkflow,omitempty"`

	// CurrentAction is the action that is currently in the running state.
	CurrentAction string `json:"currentAction,omitempty"`

	// CurrentHardware is the hardware that is currently being used.
	CurrentHardware string `json:"currentHardware,omitempty"`

	// CurrentAgent is the agent that is currently being used.
	CurrentAgent string `json:"currentAgent,omitempty"`

	// BootOptions holds the state of any boot options.
	BootOptions BootOptionsStatus `json:"bootOptions,omitempty"`

	// WorkflowRendering indicates whether the Workflow was rendered successfully.
	// Possible values are "succeeded" or "failed" or "unknown".
	WorkflowRendering WorkflowRenderingState `json:"workflowRending,omitempty"`

	// Conditions details a set of observations about the Workflow.
	// +optional
	Conditions Conditions `json:"conditions"`
}

// ActionStatus describes status information about an action.
type ActionStatus struct {
	// WorkflowID is the ID of the Workflow that the Actions belongs to.
	WorkflowID string `json:"workflowID"`

	// AgentID is the ID of the Agent that will execute the Action.
	AgentID string `json:"agentID"`

	// Rendered is the rendered action.
	Rendered Action `json:"rendered,omitempty"`

	// ID uniquely identifies the action status.
	ID string `json:"id"`

	// StartedAt is the time the action was started as reported by the client. Nil indicates the
	// Action has not started.
	StartedAt *metav1.Time `json:"startedAt,omitempty"`

	// LastTransition is the observed time when State transitioned last.
	LastTransition *metav1.Time `json:"lastTransitioned,omitempty"`

	// State describes the current state of the action.
	State ActionState `json:"state,omitempty"`

	// FailureReason is a short CamelCase word or phrase describing why the Action entered
	// ActionStateFailed.
	FailureReason string `json:"failureReason,omitempty"`

	// FailureMessage is a free-form user friendly message describing why the Action entered the
	// ActionStateFailed state. Typically, this is an elaboration on the Reason.
	FailureMessage string `json:"failureMessage,omitempty"`
}

// +kubebuilder:object:root=true

type PipelineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Pipeline `json:"items,omitempty"`
}

// BootOptions are options that control the booting of Hardware.
type BootOptions struct {
	// ToggleAllowNetboot indicates whether the controller should toggle the field in the associated hardware for allowing PXE booting.
	// This will be enabled before a Workflow is executed and disabled after the Workflow has completed successfully.
	// A HardwareRef must be provided.
	// +optional
	ToggleAllowNetboot bool `json:"toggleAllowNetboot,omitempty"`

	// ISOURL is the URL of the ISO that will be one-time booted. When this field is set, the controller will create a job.bmc.tinkerbell.org object
	// for getting the associated hardware into a CDROM booting state. A HardwareRef that contains a spec.BmcRef must be provided. If BootMode is set
	// to "isoboot", and this field is not set, a hardwareRef is required with the spec.OSIE.ISOURL field specified.
	// +optional
	// +kubebuilder:validation:Format=url
	ISOURL string `json:"isoURL,omitempty"`

	// BootMode is the type of booting that will be done. One of "netboot" or "isoboot".
	// +optional
	// +kubebuilder:validation:Enum=netboot;isoboot
	BootMode BootMode `json:"bootMode,omitempty"`
}

// BootOptionsStatus holds the state of any boot options.
type BootOptionsStatus struct {
	// AllowNetboot holds the state of the the controller's interactions with the allowPXE field in a Hardware object.
	AllowNetboot AllowNetbootStatus `json:"allowNetboot,omitempty"`
	// Jobs holds the state of any job.bmc.tinkerbell.org objects created.
	Jobs map[string]JobStatus `json:"jobs,omitempty"`
}

type AllowNetbootStatus struct {
	ToggledTrue  bool `json:"toggledTrue,omitempty"`
	ToggledFalse bool `json:"toggledFalse,omitempty"`
}

// JobStatus holds the state of a specific job.bmc.tinkerbell.org object created.
type JobStatus struct {
	// UID is the UID of the job.bmc.tinkerbell.org object associated with this workflow.
	// This is used to uniquely identify the job.bmc.tinkerbell.org object, as
	// all objects for a specific Hardware/Machine.bmc.tinkerbell.org are created with the same name.
	UID types.UID `json:"uid,omitempty"`

	// Complete indicates whether the created job.bmc.tinkerbell.org has reported its conditions as complete.
	Complete bool `json:"complete,omitempty"`

	// ExistingJobDeleted indicates whether any existing job.bmc.tinkerbell.org was deleted.
	// The name of each job.bmc.tinkerbell.org object created by the controller is the same, so only one can exist at a time.
	// Using the same name was chosen so that there is only ever 1 job.bmc.tinkerbell.org per Hardware/Machine.bmc.tinkerbell.org.
	// This makes clean up easier and we dont just orphan jobs every time.
	ExistingJobDeleted bool `json:"existingJobDeleted,omitempty"`
}

// PipelineState describes the point in time state of a Pipeline.
type PipelineState string

// ActionState describes a point in time state of an Action.
type ActionState string

// BootMode describes the type of booting that will be done.
type BootMode string

// WorkflowRenderingState describes the state of Workflow rendering.
type WorkflowRenderingState string

// PipelineConditionType is a condition used in a Pipeline. These define the state of disparate Pipeline operations.
type PipelineConditionType string

func init() {
	SchemeBuilder.Register(&Pipeline{}, &PipelineList{})
}
