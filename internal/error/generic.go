package error

import (
	"fmt"
	"os"
	"strings"

	"github.com/apex/log"
)

// CheckGenericError checks if there's an error, shows it and exits the program if it is activated
func CheckGenericError(logger *log.Entry, err error) {
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		os.Exit(1)
	}
}

// FormatGenericGolangOutput receives a string, usually the output of a command, and
// replaces the  the generic "exit status 1" that golang programs return when
// failing.
func FormatGenericGolangOutput(error string) error {
	error = strings.Replace(error, "exit status 1", "", 1)
	error = strings.Replace(error, "Error: ", "", 1)
	error = strings.ReplaceAll(error, "\n", "")
	err := fmt.Errorf("%s", error)

	return err
}
