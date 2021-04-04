package provisioner

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"gopkg.in/yaml.v2"

	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	customError "github.com/little-angry-clouds/particle/internal/error"
)

type Helm struct {
	Logger *log.Entry
}

func (h *Helm) Converge(ctx context.Context, cmd cmd.Cmd) error {
	var err error
	var logger *log.Entry = h.Logger
	var valuesKey config.Key = "values"

	values := ctx.Value(valuesKey)

	// Check correct type
	if _, ok := values.(map[string]interface{}); !ok {
		return &customError.HelmValuesType{}
	}

	// The chart is "." because it installs the chart we're developing
	// The version is "" because there's no version locally
	err = h.helmInstall(logger, cmd, ".", "", values.(map[string]interface{}))

	return err
}

func (h *Helm) Cleanup(ctx context.Context, cmd cmd.Cmd) error {
	var err error
	var name string
	var prepare config.Key = "prepare"

	c := ctx.Value(prepare)

	// Delete dependencies
	if _, ok := c.([]config.Prepare); !ok {
		return &customError.HelmPrepareType{}
	}

	for _, chart := range c.([]config.Prepare) {
		if strings.Contains(chart.Name, "/") {
			chart.Name = strings.Split(chart.Name, "/")[1]
		}
		err = h.helmDelete(cmd, chart.Name)
	}

	// Delete main chart
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	err = h.helmDelete(cmd, name)
	if err != nil {
		return err
	}

	return err
}

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

func (h *Helm) Prepare(ctx context.Context, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger
	var err error
	var prepare config.Key = "prepare"

	c := ctx.Value(prepare)

	// If there's no dependencies, or is set but not a list, just exit
	if c == nil {
		return nil
	} else if _, ok := c.([]config.Prepare); !ok {
		return &customError.HelmPrepareType{}
	}

	for _, chart := range c.([]config.Prepare) {
		// The chart is "." because it installs the chart we're developing
		err = h.helmInstall(logger, cmd, chart.Name, chart.Version, chart.Values)
	}

	return err
}


func (h *Helm) helmInstall(logger *log.Entry, cmd cmd.Cmd, chart string, version string, values map[string]interface{}) error {
	var err error
	var chartName string

	// If the chart is local, use the directory's name as chart name
	// If not, use the chart's name
	if chart == "." {
		path, err := os.Getwd()
		if err != nil {
			return err
		}

		chartName = filepath.Base(path)
	} else {
		chartName = strings.Split(chart, "/")[1]
	}

	args := []string{"helm", "upgrade", "--install", chartName, "--wait", chart}

	if version != "" {
		args = append(args, "--version")
		args = append(args, version)
	}

	// If the chart has some configuration, write it on a temporary file for helm to use it and destroy it when finished
	if values != nil {
		// Create temporary file with the defined values
		file, err := ioutil.TempFile("/tmp/", "particle-"+chartName)
		if err != nil {
			return err
		}

		defer os.Remove(file.Name())

		f, err := yaml.Marshal(&values)
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(file.Name(), f, 0644)
		if err != nil {
			return err
		}

		args = append(args, "-f", file.Name())
	}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()

	stderr := cmd.GetStderr()
	if strings.Contains(stderr, "Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused") { // nolint:lll
		err = &customError.ChartCantInstall{Name: chart}
	}

	err = customError.IsRealError(logger, err)

	return err
}

func (h *Helm) helmDelete(cmd cmd.Cmd, chart string) error {
	var logger *log.Entry = h.Logger
	var err error


	args := []string{"helm", "delete", chart}

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
