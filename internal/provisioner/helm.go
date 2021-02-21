package provisioner

import (
	"context"
	"os"
	"path/filepath"

	"github.com/little-angry-clouds/particle/internal/cmd"
)

type Helm struct{}

func (h *Helm) Converge(ctx context.Context, cmd cmd.Cmd) error {
	var err error
	var name string

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"helm", "upgrade", "--install", "test-" + name, "--wait", "."}

	err = cmd.Initialize(args)
	if err != nil {
		return err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return err
}

func (h *Helm) Cleanup(ctx context.Context, cmd cmd.Cmd) error {
	var err error
	var name string

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"helm", "delete", "test-" + name}

	err = cmd.Initialize(args)
	if err != nil {
		return err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return err
}
