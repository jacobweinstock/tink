package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/go-logr/logr"
	"github.com/go-logr/zerologr"
	"github.com/rs/zerolog"
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
	logger := getLogger(flagEnvSettings.LogLevel)

	// Randomized sleep function. Used when retrying failed connections to tink server.
	sleepFn := func() {
		rand.Seed(time.Now().UnixNano())
		s := rand.Intn(120) + 1 // 2 minutes (120 seconds), plus one second to avoid a possible zero sleep time
		logger.Info("Sleeping before attempting to re-connect...", "sleep time", s)
		time.Sleep(time.Duration(s) * time.Second)
	}

	// Here we retry failed connections to tink-server using a randomized interval between attempts.
	// tink-worker is a daemon process, so we do not exit here.
	var conn *grpc.ClientConn
	for {
		logger.Info("Obtaining tink server creds from Tink-Server URL", "Tink-Server URL", flagEnvSettings.TinkServerURL)
		creds, err := tink.ObtainServerCreds(flagEnvSettings.TinkServerURL)
		if err != nil {
			logger.Error(err, "Error obtaining server creds from Tink-Server URL", "Tink-Server URL", flagEnvSettings.TinkServerURL)
			sleepFn()
			continue
		}

		logger.Info("Connecting to tink-server...\n")
		conn, err = tink.EstablishServerConnection(flagEnvSettings.TinkServerGRPCAuthority, creds)
		if err != nil {
			logger.Error(err, "Error establishing gPRC connection to tink-server", "GRPCAuthority", flagEnvSettings.TinkServerGRPCAuthority)
			sleepFn()
			continue
		}

		break
	}

	wc := workflow.NewWorkflowServiceClient(conn)
	if wc == nil {
		logger.Error(nil, "Error creating a workflow client")
	}
}

// getLogger is a zerolog logr implementation.
func getLogger(level string) logr.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	zerologr.NameFieldName = "logger"
	zerologr.NameSeparator = "/"

	var l zerolog.Level
	switch level {
	case "debug":
		l = zerolog.DebugLevel
	case "trace":
		l = zerolog.TraceLevel
	default:
		l = zerolog.InfoLevel
	}

	zl := zerolog.New(os.Stdout).Level(l)
	zl = zl.With().Caller().Timestamp().Logger()

	return zerologr.New(&zl)
}
