package driver

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/stretchr/testify/assert"
)

type FakeCli struct {
	InitializeError error
	RunError        error
}

func (c *FakeCli) Run() error {
	return c.RunError
}

func (c *FakeCli) Initialize([]string) error {
	return c.InitializeError
}

func TestCreate(t *testing.T) { // nolint: funlen
	var test = []struct {
		testName        string
		k8sVersion      interface{}
		expectedError   error
		runError        error
		initializeError error
	}{
		{"1", config.Key("1.19.0"), nil, nil, nil},
		{"2", config.Key("1.19.0"), errors.New("initialize error"), errors.New("initialize error"), nil},
		{"3", config.Key("1.19.0"), errors.New("run error"), nil, errors.New("run error")},
		{"4", "1.19.1", errors.New("kubernetes_version has incorrect type, should be string"), nil, nil},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			var ctx context.Context = context.Background()
			var drv Kind = Kind{}
			var cli FakeCli = FakeCli{
				InitializeError: tt.initializeError,
				RunError:        tt.runError,
			}
			var kubernetesVersion config.Key = "kubernetesVersion"

			if tt.k8sVersion != "" {
				ctx = context.WithValue(ctx, kubernetesVersion, tt.k8sVersion)
			}

			err = drv.Create(ctx, &cli)
			t.Log(fmt.Sprintf("error: %s", err))
			t.Log(fmt.Sprintf("ctx value: %s", ctx.Value(kubernetesVersion)))

			assert.Equal(t, err, tt.expectedError)
		})
	}
}

func TestDestroy(t *testing.T) { // nolint: funlen
	var test = []struct {
		testName        string
		expectedError   error
		runError        error
		initializeError error
	}{
		{"1", nil, nil, nil},
		{"2", errors.New("fake error"), errors.New("fake error"), nil},
		{"3", errors.New("fake error"), nil, errors.New("fake error")},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			var ctx context.Context = context.Background()
			var drv Kind = Kind{}
			var cli FakeCli = FakeCli{
				InitializeError: tt.initializeError,
				RunError:        tt.runError,
			}

			err = drv.Destroy(ctx, &cli)
			t.Log(fmt.Sprintf("error: %s", err))

			assert.Equal(t, err, tt.expectedError)
		})
	}
}
