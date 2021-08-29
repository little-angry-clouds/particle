package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/apex/log"
)

// A Cmd defines an interface to execute some operations with os/exec.
// Is composed by:
//   - Initialize: It will do what's necessary to be able to run a command, like set the stdout or the logger.
//   - Run: It will run the command, duh.
//   - GetStderr: It will return the stderr of the command. It's usually used to determinate if an error is really an error in particle.
type Cmd interface {
	Initialize(*log.Entry, []string) error
	Run() error
	GetStderr() string
}

// A CLI defines the structure that will be used for the Cmd interface.
// It's composed by:
//   - Binary: The name of the binary that will be executed.
//   - Path: The path of the binary that will be executed.
//   - Args: A list of strings that will be passed as arguments to the command executed.
//   - Stderr: An interface that will be used to write the stderr.
//   - Stdout: An interface that will be used to write the stdout.
//   - Logger: The apex log entry.
//   - stderrString: The stderr in a string format, used only internally.
type CLI struct {
	Binary       string
	Path         string
	Args         []string
	Stderr       io.Writer
	Stdout       io.Writer
	Logger       *log.Entry
	stderrString string
}

// Initialize sets the struct attributes necessary to execute the command.
func (c *CLI) Initialize(logger *log.Entry, args []string) error {
	c.Args = args
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout
	c.Logger = logger

	path, err := exec.LookPath(c.Binary)
	if err != nil {
		return err
	}

	c.Path = path

	return nil
}

// Run executes the command. If debug is set, the stderr and stdout of the executed
// command will be actually redirected to the terminals stderr and stdout.
// If not, it will be written in a variable used to further analysis of the errors.
func (c *CLI) Run() error {
	var err error
	var stdout, stderr bytes.Buffer

	cmd := &exec.Cmd{}
	cmd.Args = c.Args
	cmd.Path = c.Path

	if c.Logger.Level == log.DebugLevel {
		cmd.Stderr = io.MultiWriter(c.Stderr, &stderr)
		cmd.Stdout = io.MultiWriter(c.Stdout, &stdout)
	} else {
		cmd.Stderr = &stderr
		cmd.Stdout = &stdout
	}

	c.Logger.Debug(
		fmt.Sprintf(
			"Command to execute: %s",
			strings.Replace(
				strings.Join(cmd.Args, " "),
				"\n",
				" && ",
				-1),
		),
	)

	err = cmd.Run()

	c.stderrString = stderr.String()

	return err
}

// GetStderr returns the stderr in string format.
func (c *CLI) GetStderr() string {
	return c.stderrString
}
