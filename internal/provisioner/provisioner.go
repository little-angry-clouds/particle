package provisioner

import (
	"context"

	"github.com/little-angry-clouds/particle/internal/cmd"
)

type Provisioner interface {
	Converge(context.Context, cmd.Cmd) error
}
