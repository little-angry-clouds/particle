package provisioner

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"gopkg.in/yaml.v2"

	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

type Helm struct {
	Logger *log.Entry
}

func (h *Helm) Converge(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var err error
	var logger *log.Entry = h.Logger
	var values map[string]interface{} = configuration.Provisioner.Values

	err = h.helmInstall(logger, cmd, ".", "", values)

	return err
}

func (h *Helm) Cleanup(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger
	var err error
	var name string
	var prepare []config.Prepare = configuration.Prepare

	for _, chart := range prepare {
		if strings.Contains(chart.Name, "/") {
			chart.Name = strings.Split(chart.Name, "/")[1]
		}

		err = h.helmDelete(cmd, chart.Name)
		err = isRealError(logger, err)

		if err != nil {
			return err
		}
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

func (h *Helm) Dependency(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger
	var err error
	var args []string
	var dependencies []config.Dependencies = configuration.Dependency.Charts

	// If there's no dependencies, or is set but not a list, just exit
	if dependencies == nil {
		return nil
	}

	for _, dependency := range dependencies {
		args = []string{"helm", "repo", "add", dependency.RepositoryName, dependency.RepositoryURL}

		err = cmd.Initialize(logger, args)
		if err != nil {
			return err
		}

		err = cmd.Run()

		stderr := cmd.GetStderr()
		if strings.Contains(stderr, "already exists, please specify a different name") {
			err = &helmRepoExists{Name: dependency.RepositoryName}
		}

		err = isRealError(logger, err)
		if err != nil {
			return err
		}
	}

	return err
}

func (h *Helm) Prepare(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger
	var err error
	var dependencies []config.Prepare = configuration.Prepare

	// If there's no dependencies, or is set but not a list, just exit
	if dependencies == nil {
		return nil
	}

	for _, chart := range dependencies {
		err = h.helmInstall(logger, cmd, chart.Name, chart.Version, chart.Values)
	}

	return err
}

func (h *Helm) helmInstall(logger *log.Entry, cmd cmd.Cmd, chart string, version string, values map[string]interface{}) error {
	var err error
	var chartName string

	// If the chart is local, use the directory's name as chart name
	// If not, use the chart's name
	// The chart is "." because it installs the chart we're developing
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
		args = append(args, "--version", version)
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
	if strings.Contains(stderr, "Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused") {
		err = &chartCantInstall{Name: chart}
	}

	return isRealError(logger, err)
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
		err = &chartNotInstalled{Name: chart}
	} else if strings.Contains(stderr, "Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused") {
		err = &chartCantDelete{Name: chart}
	}

	err = isRealError(logger, err)

	return err
}
