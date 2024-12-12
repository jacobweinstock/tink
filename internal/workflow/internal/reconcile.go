package internal

import (
	"bytes"
	"context"
	"text/template"
	"time"

	"github.com/go-logr/logr"
	"github.com/google/uuid"
	tinkv1 "github.com/tinkerbell/tink/api/v1alpha2"
	"gopkg.in/yaml.v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// ReconciliationContext reconciles Workflow resources when created or updated.
type ReconciliationContext struct {
	// Pipeline is the Pipeline instance we're reconciling.
	Pipeline *tinkv1.Pipeline

	// NewActionID generated unique IDs for actions. Defaults to generating UUIDv4s.
	NewActionID func() string

	Log    logr.Logger
	Client client.Client
}

// Reconcile reconciles the Workflow.
func (rc ReconciliationContext) Reconcile(ctx context.Context) (reconcile.Result, error) {
	wfRef := client.ObjectKey{
		Name:      rc.Pipeline.Spec.Workflows[0].WorkflowRef.Name,
		Namespace: rc.Pipeline.Namespace,
	}
	var tmpl tinkv1.Workflow
	if err := rc.Client.Get(ctx, wfRef, &tmpl); err != nil {
		if errors.IsNotFound(err) {
			// The Template may yet to be submitted to the cluster so just requeue.
			rc.Log.Info("Template not found; requeue in 5 seconds", "ref", wfRef)
			return reconcile.Result{RequeueAfter: 5 * time.Second}, nil
		}
		return reconcile.Result{}, err
	}

	hwRef := client.ObjectKey{
		Name:      rc.Pipeline.Spec.Workflows[0].HardwareRef.Name,
		Namespace: rc.Pipeline.Namespace,
	}
	var hw tinkv1.Hardware
	if err := rc.Client.Get(ctx, hwRef, &hw); err != nil {
		if errors.IsNotFound(err) {
			// The Hardware may yet to be submitted to the cluster so just requeue.
			rc.Log.Info("Hardware not found; requeue in 5 seconds", "ref", wfRef)
			return reconcile.Result{RequeueAfter: 5 * time.Second}, nil
		}
		return reconcile.Result{}, err
	}

	// Only render the template and configure action status if its not been done before.
	if len(rc.Pipeline.Status.Workflows) == 0 {
		tmpl, err := rc.renderTemplate(tmpl, &hw)
		if err != nil {
			return reconcile.Result{}, err
		}

		rc.Pipeline.Status.Workflows = rc.toActionStatus(tmpl.Spec.Actions)
	}

	return reconcile.Result{}, nil
}

func (rc ReconciliationContext) renderTemplate(tpl tinkv1.Workflow, hw *tinkv1.Hardware) (tinkv1.Workflow, error) {
	tplYAML, err := yaml.Marshal(tpl)
	if err != nil {
		return tinkv1.Workflow{}, err
	}

	renderer, err := template.New("").
		Option("missingkey=error").
		Funcs(workflowTemplateFuncs).
		Parse(string(tplYAML))
	if err != nil {
		return tinkv1.Workflow{}, err
	}

	tplData := map[string]any{
		"Hardware": hw.Spec,
		"Param":    rc.Pipeline.Spec.TemplateParams,
	}

	var renderedTplYAML bytes.Buffer
	if err := renderer.Execute(&renderedTplYAML, tplData); err != nil {
		return tinkv1.Workflow{}, err
	}

	if err := yaml.Unmarshal(renderedTplYAML.Bytes(), &tpl); err != nil {
		return tinkv1.Workflow{}, err
	}

	return tpl, nil
}

func (rc ReconciliationContext) toActionStatus(actions []tinkv1.Action) []tinkv1.ActionStatus {
	var status []tinkv1.ActionStatus
	for _, action := range actions {
		status = append(status, tinkv1.ActionStatus{
			Rendered: action,
			ID:       rc.newActionID(),
			State:    tinkv1.ActionStatePending,
		})
	}
	return status
}

func (rc ReconciliationContext) newActionID() string {
	if rc.NewActionID != nil {
		return rc.NewActionID()
	}
	return uuid.New().String()
}
