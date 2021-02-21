package verifier

import (
	"context"
	"os"
	"path/filepath"

	"github.com/little-angry-clouds/particle/internal/cmd"
)

type Helm struct{}

func (h *Helm) Verify(ctx context.Context, cmd cmd.Cmd) error {
	var err error
	var name string

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"helm", "test", "test-" + name}

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
