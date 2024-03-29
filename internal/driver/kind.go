package driver

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	"gopkg.in/yaml.v2"
)

// Kind is an implementation of the driver interface. It uses kind to manage the kubernetes cluster:
// https://kind.sigs.k8s.io/
type Kind struct {
	Logger *log.Entry
}

// Create creates the kubernetes cluster. It will need the k8s cluster version, which is in the
// configuration, and the name will be created from the path of execution, to give it a more unique name.
// It's basicallt a wrap for the "kind" command. All of its configuration can be don trough particle's
// configuration.
func (k *Kind) Create(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = k.Logger
	var kubernetesVersion string = configuration.Driver.KubernetesVersion
	var values map[string]interface{} = configuration.Driver.Values
	var err error
	var name string

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"kind", "create", "cluster", "--wait", "1m", "--name", name}

	// Check if k8s version is set. If not, use latest.
	if kubernetesVersion != "" {
		args = append(args, []string{"--image", "kindest/node:" + kubernetesVersion}...)
	}

	// If kind has some configuration, write it on a temporary file for kind to use it and destroy it when finished
	if values != nil {
		// Create temporary file with the defined values
		file, err := ioutil.TempFile("/tmp/", "particle-kind-"+name)
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

		args = append(args, fmt.Sprintf("--config=%s", file.Name()))
	}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()

	stderr := cmd.GetStderr()
	if strings.Contains(stderr, "failed to create cluster: node(s) already exist for a cluster with the name") {
		err = &clusterExists{Name: name}
	}

	err = isRealError(logger, err)

	return err
}

// Destroy destroys the kubernetes cluster. Not much to add to it.
func (k *Kind) Destroy(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = k.Logger
	var err error
	var name string

	// Use the directory as the cluster ID
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"kind", "delete", "cluster", "--name", name}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()

	return err
}
