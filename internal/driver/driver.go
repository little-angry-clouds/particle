package driver

import (
	"context"

	"github.com/little-angry-clouds/particle/internal/cmd"
)

type Driver interface {
	Create(context.Context, cmd.Cmd) error
	Destroy(context.Context, cmd.Cmd) error
}
