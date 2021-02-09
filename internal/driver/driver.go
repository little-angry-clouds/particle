package driver

import (
	"context"
)

type Driver interface {
	Create(context.Context, Cmd) error
	Destroy(context.Context) error
}

