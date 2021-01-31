package helpers

import (
	"fmt"
	"os"
)

// CheckGenericError checks if there's an error, shows it and exits the program if it is
func CheckGenericError(err error) {
	if err != nil {
		message := fmt.Sprintf("An error was detected, exiting: %s", err)
		fmt.Fprintf(os.Stderr, "%s\n", message)
		os.Exit(1) // nolint:gomnd
	}
}

// StringInSlice checks if there's a string in a slice
func StringInSlice(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
