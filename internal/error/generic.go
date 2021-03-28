package error

import (
	"os"
	"strings"

	"github.com/apex/log"
)

// CheckGenericError checks if there's an error, shows it and exits the program if it is activated
func CheckGenericError(logger *log.Entry, err error, exit bool) {
	if err != nil {
		switch err.(type) {
		// Do nothing if this error, since always exists when there's no cluster and cleanup is executed
		case *InexistentCluster:
		default:
			logger.WithError(err).Error("An error was detected, exiting")

			if exit {
				os.Exit(1)
			}
		}
	}
}

// ManageError returns custom error if found, if not returns the generic.
func ManageError(err error, output string) error {
	var e error

	switch {
	// This output is returned by kind when there's an existing cluster
	case strings.Contains(output, "failed to create cluster: node(s) already exist for a cluster with the name"):
		e = &ClusterExists{}
	// This output is returned by helm cleanup when there's no existing cluster
	case strings.Contains(output, "Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused"): // nolint:lll
		e = &InexistentCluster{}
	default:
		e = err
	}

	return e
}
