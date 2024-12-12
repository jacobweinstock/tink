package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type WorkflowSpec struct {
	// Actions defines the set of Actions to be run by an Agent. Actions are run sequentially
	// in the order they are specified. At least 1 Action must be specified. Names of Actions
	// must be unique within a Workflow.
	// +kubebuilder:validation:MinItems=1
	Actions []Action `json:"actions,omitempty"`

	// Volumes to be mounted on all Actions. If an Action specifies the same volume it will take
	// precedence.
	// +optional
	Volumes []Volume `json:"volumes,omitempty"`

	// Env defines environment variables to be available in all Actions. If an Action specifies
	// the same environment variable, the Action's environment variable will take precedence.
	// +optional
	Env map[string]string `json:"env,omitempty"`

	// Logging defines the logging configuration for all Action containers. If not specified the
	// default runtime configured logging driver is used. This is Agent and runtime specific.
	// +optional
	Logging *Logging `json:"logging,omitempty"`
}

// Action defines an individual Action to be run on a target machine.
type Action struct {
	// Name is a name for the Action.
	Name string `json:"name"`

	// Image is fully qualified OCI image name.
	Image string `json:"image"`

	// Cmd defines the command to use when launching the image. It overrides the default command
	// defined in the Image.
	// +optional
	Cmd *string `json:"cmd,omitempty"`

	// Args are a set of arguments to be passed to the command executed by the container on launch.
	// +optional
	Args []string `json:"args,omitempty"`

	// Env defines environment variables used when launching the container.
	//+optional
	Env map[string]string `json:"env,omitempty"`

	// Volumes defines the volumes to mount into the container.
	// +optional
	Volumes []Volume `json:"volumes,omitempty"`

	// Namespace defines the Linux namespaces this container should execute in.
	// +optional
	Namespace *Namespace `json:"namespaces,omitempty"`

	// TimeoutSeconds defines the time the Action has to complete. The timer begins when the action is requested.
	// When set to 0, no timeout is applied.
	// +kubebuilder:default=0
	// +kubebuilder:validation:Minimum=0
	TimeoutSeconds *int64 `json:"timeout,omitempty"`

	// Background true will run the container in the background or as some runtimes call it detached mode. When set to true, the Agent will immediately report the Action as having succeeded.
	// +optional
	Background *bool `json:"background,omitempty"`

	// Retries defines the number of times the Agent will retry the Action if it fails, not including Timeouts. Retries are only attempted on non-zero exit codes.
	// +optional
	Retries *int32 `json:"retries,omitempty"`
}

// Volume is a specification for mounting a volume in an Action. Volumes take the form
// {SRC-VOLUME-NAME | SRC-HOST-DIR}:TGT-CONTAINER-DIR:OPTIONS. When specifying a VOLUME-NAME that
// does not exist it will be created for you. Examples:
//
// Read-only bind mount bound to /data
//
//	/etc/data:/data:ro
//
// Writable volume name bound to /data
//
//	shared_volume:/data
//
// See https://docs.docker.com/storage/volumes/ for additional details.
type Volume string

// Namespace defines the Linux namespaces to use for the container.
// See https://man7.org/linux/man-pages/man7/namespaces.7.html.
type Namespace struct {
	// Network defines the network namespace.
	// +optional
	Network *string `json:"network,omitempty"`

	// PID defines the PID namespace
	// +optional
	PID *string `json:"pid,omitempty"`
}

// Logging defines
type Logging struct {
	// Driver is the logging driver to use for Action containers.
	Driver string `json:"driver,omitempty"`

	// Options are the logging options to use for Action containers.
	Options map[string]string `json:"options,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:categories=tinkerbell,shortName=bp

// Workflow defines a set of Actions to be run by an Agent. The Workflow is rendered
// prior to execution where it is exposed to Hardware and user defined data. Most fields within the
// WorkflowSpec may contain template values excluding .Workflow.Spec.Actions[].Name.
// See https://pkg.go.dev/text/template for more details.
type Workflow struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WorkflowSpec `json:"spec,omitempty"`
}

// +kubebuilder:object:root=true

type WorkflowList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workflow `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workflow{}, &WorkflowList{})
}
