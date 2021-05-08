package provisioner

import (
	"fmt"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/helpers"
)

type clusterNotExists struct{}

func (e *clusterNotExists) Error() string {
	return "there's no cluster detected"
}

type helmRepoExists struct {
	Name string
}

func (e *helmRepoExists) Error() string {
	return fmt.Sprintf("the helm repository '%s' is already added", e.Name)
}

type chartNotInstalled struct {
	Name string
}

func (e *chartNotInstalled) Error() string {
	return fmt.Sprintf("chart '%s' is not installed", e.Name)
}

type chartCantInstall struct {
	Name string
}

func (e *chartCantInstall) Error() string {
	return fmt.Sprintf("%s, so '%s' can't be installed", &clusterNotExists{}, e.Name)
}

type chartCantDelete struct {
	Name string
}

func (e *chartCantDelete) Error() string {
	return fmt.Sprintf("%s, so '%s' can't be deleted", &clusterNotExists{}, e.Name)
}

func isRealError(logger *log.Entry, err error) error {
	switch err.(type) {
	// exists when an helm repository is already added
	case *helmRepoExists:
		logger.Warn(helpers.Capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// exists when a chart is not installed
	case *chartNotInstalled:
		logger.Warn(helpers.Capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// exists when the chart can't be deleted
	case *chartCantDelete:
		logger.Warn(helpers.Capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// exists when the chart can't be installed
	case *chartCantInstall:
		logger.Warn(helpers.Capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// Return received error
	default:
	}

	return err
}
