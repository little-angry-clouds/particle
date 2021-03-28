package cmd

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/apex/log"

	customError "github.com/little-angry-clouds/particle/internal/error"
)

type Cmd interface {
	Initialize(*log.Entry, []string) error
	Run() error
}

type CLI struct {
	Binary string
	Path   string
	Args   []string
	Stderr io.Writer
	Stdout io.Writer
	Logger *log.Entry
}

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

	err = customError.ManageError(cmd.Run(), stderr.String())

	return err
}

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
