package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/tinkerbell/tink/cmd/tink-worker/client/tink"
	"github.com/tinkerbell/tink/cmd/tink-worker/cmd"
	"github.com/tinkerbell/tink/protos/workflow"
	"google.golang.org/grpc"
)

func main() {
	// parse and validate command-line flags and required env vars
	flagEnvSettings, err := cmd.CollectFlagEnvSettings(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	// Randomized sleep function. Used when retrying failed connections to tink server.
	sleepFn := func() {
		rand.Seed(time.Now().UnixNano())
		s := rand.Intn(120) + 1 // 2 minutes (120 seconds), plus one second to avoid a possible zero sleep time
		fmt.Fprintf(os.Stderr, "Sleeping %d seconds before attempting to re-connect...\n", s)
		time.Sleep(time.Duration(s) * time.Second)
	}

	// Here we retry failed connections to tink-server using a randomized interval between attempts.
	// tink-worker is a daemon process, so we do not exit here.
	var conn *grpc.ClientConn
	for {
		fmt.Fprintf(os.Stderr, "Obtaining tink server creds from %s...\n", flagEnvSettings.TinkServerURL)
		creds, err := tink.ObtainServerCreds(flagEnvSettings.TinkServerURL)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error obtaining server creds from %s: %v\n", flagEnvSettings.TinkServerURL, err)
			sleepFn()
			continue
		}

		fmt.Fprintf(os.Stderr, "Connecting to tink-server...\n")
		conn, err = tink.EstablishServerConnection(flagEnvSettings.TinkServerGRPCAuthority, creds)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error establishing gPRC connection to tink-server at %s: %v\n", flagEnvSettings.TinkServerGRPCAuthority, err)
			sleepFn()
			continue
		}

		break
	}

	wc := workflow.NewWorkflowServiceClient(conn)
	if wc == nil {
		fmt.Fprintf(os.Stderr, "Error creating a workflow client")
	}
}
