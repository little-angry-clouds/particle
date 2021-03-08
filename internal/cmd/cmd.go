package cmd

import (
	"io"
	"os"
	"os/exec"
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
	cmd := &exec.Cmd{}
	cmd.Args = c.Args
	cmd.Path = c.Path
	cmd.Stderr = c.Stderr
	cmd.Stdout = c.Stdout

	return cmd.Run()
}

func (c *CLI) Initialize(args []string) error {
	c.Args = args
	c.Stderr = os.Stderr
	c.Stdout = os.Stdout

	path, err := exec.LookPath(c.Binary)
	if err != nil {
		return err
	}

	c.Path = path

	return nil
}
