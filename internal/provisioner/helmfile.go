package provisioner

import (
	"strings"

	"github.com/apex/log"

	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
	customError "github.com/little-angry-clouds/particle/internal/error"
)

// Helmfile is an implementation of the provisioner interface. It uses helm to manage the kubernetes cluster:
// https://github.com/roboll/helmfile/
type Helmfile struct {
	Logger *log.Entry
}

const debug string = "debug"

// Converge ensures that the deployment is executed on the kubernetes cluster.
func (h *Helmfile) Converge(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var err error
	var logger *log.Entry = h.Logger
	var args []string

	if string(rune(logger.Level)) == debug {
		args = []string{"helmfile", "--debug", "sync"}
	} else {
		args = []string{"helmfile", "sync"}
	}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	stderr := cmd.GetStderr()
	if strings.Contains(stderr, "Kubernetes cluster unreachable: Get") {
		err = &clusterUnreachable{}
	} else if strings.Contains(stderr, "exit status 1") {
		err = customError.FormatGenericGolangOutput(stderr)
	}

	return isRealError(logger, err)
}

// Prepare does nothing, since helmfile can install multiple charts in the order you need, so it's unnecessary.
func (h *Helmfile) Cleanup(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger
	var err error
	var args []string

	if string(rune(logger.Level)) == debug {
		args = []string{"helmfile", "--debug", "destroy"}
	} else {
		args = []string{"helmfile", "destroy"}
	}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()

	stderr := cmd.GetStderr()

	switch {
	case strings.Contains(stderr, "Kubernetes cluster unreachable: Get"):
		err = &clusterUnreachable{}
	case strings.Contains(stderr, "exit status 1"):
		err = customError.FormatGenericGolangOutput(stderr)
	}

	return isRealError(logger, err)
}

// Dependency locally adds all the helm repositories. It basically executes "helm repo add $whatever" under the helmfile umbrella.
// Also, it's pretty useless since helmfile sync, which is what Converge executes, also executes it. But it's added for debugging purposes, if needed.
func (h *Helmfile) Dependency(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var err error
	var logger *log.Entry = h.Logger
	var args []string

	if string(rune(logger.Level)) == "debug" {
		args = []string{"helmfile", "--debug", "repos"}
	} else {
		args = []string{"helmfile", "repos"}
	}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()

	return err
}

// Prepare does nothing, since helmfile can install multiple charts in the order you need, so it's unnecessary.
func (h *Helmfile) Prepare(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = h.Logger

	logger.Warn("This step is skipped, since you can do whatever preparation you need with the helm provider.")

	return nil
}
