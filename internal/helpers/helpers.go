package helpers

import (
	"os"
	"unicode"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"
)

// CheckGenericError checks if there's an error, shows it and exits the program if it is
func CheckGenericError(logger *log.Entry, err error) {
	if err != nil {
		logger.WithError(err).Error("An error was detected, exiting")
		os.Exit(1)
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

// Set the logger level
func GetLogger(debug bool) *log.Entry {
	var logger log.Entry

	if debug {
		logger = log.Entry{
			Logger: &log.Logger{
				Handler: cli.Default,
				Level:   log.DebugLevel,
			},
			Level: log.DebugLevel,
		}
	} else {
		logger = log.Entry{
			Logger: &log.Logger{
				Handler: cli.Default,
				Level:   log.InfoLevel,
			},
			Level: log.InfoLevel,
		}
	}

	return &logger
}

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}
