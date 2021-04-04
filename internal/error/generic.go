package error

import (
	"fmt"
	"os"
	"unicode"

	"github.com/apex/log"
)

// CheckGenericError checks if there's an error, shows it and exits the program if it is activated
func CheckGenericError(logger *log.Entry, err error) {
	err = IsRealError(logger, err)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		os.Exit(1)
	}
}

// isRealError returns true or false if it's a real error
func IsRealError(logger *log.Entry, err error) error {
	switch err.(type) {
	// exists when there's a cluster created
	case *ClusterExists:
		logger.Warn(capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// exists when an helm repository is already added
	case *HelmRepoExists:
		logger.Warn(capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// exists when a chart is not installed
	case *ChartNotInstalled:
		logger.Warn(capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// exists when the chart can't be deleted
	case *ChartCantDelete:
		logger.Warn(capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// exists when the chart can't be installed
	case *ChartCantInstall:
		logger.Warn(capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// Return received error
	default:
	}

	return err
}

func capitalize(s string) string {
    if len(s)==0 {
       return s
    }

    r := []rune(s)
    r[0] = unicode.ToUpper(r[0])
    return string(r)
}
