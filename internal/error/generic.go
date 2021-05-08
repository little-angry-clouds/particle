package error

import (
	"os"

	"github.com/apex/log"
)

// CheckGenericError checks if there's an error, shows it and exits the program if it is activated
func CheckGenericError(logger *log.Entry, err error) {
	// err = IsRealError(logger, err)
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		os.Exit(1)
	}
}
