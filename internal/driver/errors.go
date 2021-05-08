package driver

import (
	"fmt"

	"github.com/apex/log"

	"github.com/little-angry-clouds/particle/internal/helpers"
)

type clusterExists struct {
	Name string
}

func (e *clusterExists) Error() string {
	return fmt.Sprintf("cluster '%s' already exists", e.Name)
}

func isRealError(logger *log.Entry, err error) error {
	switch err.(type) {
	// exists when an cluster is already created
	case *clusterExists:
		logger.Warn(helpers.Capitalize(fmt.Sprintf("%s", err)))
		err = nil
	// Return received error
	default:
	}

	return err
}
