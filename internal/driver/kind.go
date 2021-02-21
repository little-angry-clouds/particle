package driver

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

type Kind struct{}

func (k *Kind) Create(ctx context.Context, cmd cmd.Cmd) error {
	var err error
	var name string
	var kubernetesVersion config.Key = "kubernetesVersion"

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"kind", "create", "cluster", "--wait", "1m", "--name", name}

	// Check if k8s version is set
	version := ctx.Value(kubernetesVersion)
	if version != nil {
		// If k8s version is set, check it's a string
		if value, ok := version.(config.Key); ok {
			args = append(args, []string{"--image", "kindest/node:" + string(value)}...)
		} else {
			return errors.New("kubernetes_version has incorrect type, should be string")
		}
	}

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

func (k *Kind) Destroy(ctx context.Context, cmd cmd.Cmd) error {
	var err error
	var name string

	// Use the directory as the cluster ID
	path, err := os.Getwd()
	if err != nil {
		return err
	}

	name = filepath.Base(path)

	args := []string{"kind", "delete", "cluster", "--name", name}

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
