package driver

import (
	"context"
	"errors"
	"fmt"
	"io"
	"testing"

	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/stretchr/testify/assert"
)

type FakeCli struct {
	Path   string
	Args   []string
	Stderr io.Writer
}

func (c *FakeCli) Run() error {
	return nil
}

func (c *FakeCli) Initialize([]string) error {
	return nil
}

func TestCreate(t *testing.T) { // nolint: funlen
	var test = []struct {
		testName   string
		k8sVersion interface{}
		error      error
	}{
		{"1", config.Key("1.19.0"), nil},
		{"2", "1.19.1", errors.New("kubernetes_version has incorrect type, should be string")},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			var ctx context.Context = context.Background()
			var drv Kind = Kind{}
			var cli FakeCli = FakeCli{}
			var kubernetesVersion config.Key = "kubernetesVersion"

	if tt.k8sVersion != "" {
			ctx = context.WithValue(ctx, kubernetesVersion, tt.k8sVersion)
	}
			err = drv.Create(ctx, &cli)
			t.Log(fmt.Sprintf("error: %s", err))
			t.Log(fmt.Sprintf("ctx value: %s", ctx.Value(kubernetesVersion)))

			assert.Equal(t, err, tt.error)
		})
	}
}
