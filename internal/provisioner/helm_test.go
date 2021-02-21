package provisioner

import (
	"context"
	"errors"
	"fmt"
	"testing"

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

func TestConverge(t *testing.T) { // nolint: funlen
	var test = []struct {
		testName        string
		expectedError   error
		runError        error
		initializeError error
	}{
		{"1", nil, nil, nil},
		{"2", errors.New("initialize error"), errors.New("initialize error"), nil},
		{"3", errors.New("run error"), nil, errors.New("run error")},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			var ctx context.Context = context.Background()
			var drv Helm = Helm{}
			var cli FakeCli = FakeCli{
				InitializeError: tt.initializeError,
				RunError:        tt.runError,
			}

			err = drv.Converge(ctx, &cli)
			t.Log(fmt.Sprintf("error: %s", err))

			assert.Equal(t, err, tt.expectedError)
		})
	}
}
