package worker

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/docker/docker/client"
	"github.com/go-logr/stdr"
	"github.com/tinkerbell/tink/cmd/tink-worker/container"
	"github.com/tinkerbell/tink/protos/workflow"
)

func TestController(t *testing.T) {
	t.Skip()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	logger := stdr.New(log.New(os.Stdout, "", log.Lshortfile))
	queue := make(chan *workflow.WorkflowAction)
	go actionController(ctx, logger, queue)
	queue <- &workflow.WorkflowAction{
		TaskName: "myTask",
		Name:     "actionOne",
		Image:    "alpine",
		Timeout:  0,
		Command:  []string{"echo", "hello"},
		WorkerId: "123",
	}
	queue <- &workflow.WorkflowAction{
		TaskName: "myTask",
		Name:     "actionTwo",
		Image:    "alpine",
		Timeout:  0,
		Command:  []string{"echo", "hello"},
		WorkerId: "123",
	}
	time.Sleep(time.Second * 6)
	t.Fail()
}

func TestExecuteAction(t *testing.T) {
	t.Skip()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	logger := stdr.New(log.New(os.Stdout, "", log.Lshortfile|log.Ltime))
	action := &workflow.WorkflowAction{
		TaskName: "myTask",
		Name:     "actionOne",
		Image:    "golang",
		Command:  []string{"bash", "-c", "cd /boots; make boots"},
		Volumes:  []string{"/home/tink/repos/tinkerbell/boots:/boots"},
		WorkerId: "123",
	}
	conn, err := client.NewClientWithOpts()
	if err != nil {
		logger.Error(err, "failed to create docker client")
		return
	}
	cm := &container.Docker{Conn: conn}

	Run(ctx, logger, cm, action)
	t.Fail()
}
