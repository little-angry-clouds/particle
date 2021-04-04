package driver

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	customError "github.com/little-angry-clouds/particle/internal/error"
)

type Kind struct {
	Logger *log.Entry
}

func (k *Kind) Create(ctx context.Context, cmd cmd.Cmd) error {
	var logger *log.Entry = k.Logger
	var kubernetesVersion config.Key = "kubernetesVersion"
	var err error
	var name string

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
			return &customError.KubernetesVersionType{}
		}
	}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()

	stderr := cmd.GetStderr()
	if strings.Contains(stderr, "failed to create cluster: node(s) already exist for a cluster with the name") {
		err = &customError.ClusterExists{Name: name}
	}

	err = customError.IsRealError(logger, err)

	return err
}

func (k *Kind) Destroy(ctx context.Context, cmd cmd.Cmd) error {
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
