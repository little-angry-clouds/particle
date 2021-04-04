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
func (h *Helm) Dependency(ctx context.Context, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger
	var err error
	var charts config.Key = "charts"
	var args []string

	dependencies := ctx.Value(charts)

	// If there's no dependencies, or is set but not a list, just exit
	if dependencies == nil {
		return nil
	} else if _, ok := dependencies.([]config.Dependencies); !ok {
		return &customError.HelmDependencyType{}
	}

	for _, dependency := range dependencies.([]config.Dependencies) {
		args = []string{"helm", "repo", "add", dependency.RepositoryName, dependency.RepositoryURL}

		err = cmd.Initialize(logger, args)
		if err != nil {
			return err
		}

		err = cmd.Run()

		stderr := cmd.GetStderr()
		if strings.Contains(stderr, "already exists, please specify a different name") {
			err = &customError.HelmRepoExists{Name: dependency.RepositoryName}
		}

		err = customError.IsRealError(logger, err)
		if err != nil {
			return err
		}
	}

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
	stderr := cmd.GetStderr()
	if strings.Contains(stderr, "Release not loaded") {
		err = &customError.ChartNotInstalled{Name: chart}
	} else if strings.Contains(stderr, "Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused") { // nolint:lll
		err = &customError.ChartCantDelete{Name: chart}
	}

	err = customError.IsRealError(logger, err)

	return err
}
