package verifier

import (
	"context"

	"github.com/little-angry-clouds/particle/internal/cmd"
)

type Verifier interface {
	Verify(context.Context, cmd.Cmd) error
}
