package driver

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

type Kind struct {
	Logger *log.Entry
}

func (k *Kind) Create(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = k.Logger
	var kubernetesVersion string = configuration.Driver.KubernetesVersion
	var err error
	var name string

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"kind", "create", "cluster", "--wait", "1m", "--name", name}

	// Check if k8s version is set
	if kubernetesVersion != "" {
		args = append(args, []string{"--image", "kindest/node:" + kubernetesVersion}...)
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
