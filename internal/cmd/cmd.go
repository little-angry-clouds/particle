package cmd

import (
	"io"
	"os"
	"os/exec"
)

type Cmd interface {
	Initialize([]string) error
	Run() error
}

type CLI struct {
	Binary string
	Path   string
	Args   []string
	Stderr io.Writer
}

func (c *CLI) Run() error {
	cmd := &exec.Cmd{}
	cmd.Args = c.Args
	cmd.Path = c.Path
	cmd.Stderr = c.Stderr

	return cmd.Run()
}

func (c *CLI) Initialize(args []string) error {
	c.Args = args
	c.Stderr = os.Stderr

	path, err := exec.LookPath(c.Binary)
	if err != nil {
		return err
	}

	c.Path = path

	return nil
}
