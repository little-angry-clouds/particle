package provisioner

import (
	"context"
	"os"
	"path/filepath"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
)

type Helm struct {
	Logger *log.Entry
}

func (h *Helm) Converge(ctx context.Context, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger
	var err error
	var name string

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"helm", "upgrade", "--install", "test-" + name, "--wait", "."}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()

	return err
}

func (h *Helm) Cleanup(ctx context.Context, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger
	var err error
	var name string

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"helm", "delete", "test-" + name}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()

	return err
}
