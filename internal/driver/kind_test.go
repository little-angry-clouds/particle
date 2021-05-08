package driver

import (
	"errors"
	"fmt"
	"testing"

	"github.com/apex/log"

	"github.com/little-angry-clouds/particle/internal/config"
	"github.com/little-angry-clouds/particle/internal/helpers"
	"github.com/stretchr/testify/assert"
)

type FakeCli struct {
	CliError error
	Stderr   string
}

func (c *FakeCli) Run() error {
	return c.CliError
}

func (c *FakeCli) Initialize(*log.Entry, []string) error {
	return nil
}

func (c *FakeCli) GetStderr() string {
	return c.Stderr
}

func TestCreate(t *testing.T) {
	var test = []struct {
		testName      string
		expectedError error
		cliError      error
		stdErr        string
		k8sVersion    string
	}{
		// Test that the create function works with no error
		{"1", nil, nil, "", ""},
		// Test that the create function works with no error
		{"1", nil, nil, "", "1.19.0"},
		// Test that unexpected generic error is handled as an error
		{"2",
			// expectedError
			errors.New("fake error"),
			// cliError
			errors.New("fake error"),
			// stdErr
			"",
			// k8sVersion
			"",
		},
		// Test that clusterExists error is not handled as an error
		{"3",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"failed to create cluster: node(s) already exist for a cluster with the name",
			// k8sVersion
			"",
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			var drv Kind = Kind{Logger: helpers.GetLogger(false)}
			var cli FakeCli = FakeCli{
				CliError: tt.cliError,
				Stderr:   tt.stdErr,
			}
			var configuration config.ParticleConfiguration

			configuration.Driver.KubernetesVersion = tt.k8sVersion
			err = drv.Create(configuration, &cli)
			t.Log(fmt.Sprintf("error: %s", err))
			t.Log(fmt.Sprintf("config value: %s", configuration))

			assert.Equal(t, err, tt.expectedError)
		})
	}
}

func TestDestroy(t *testing.T) {
	var test = []struct {
		testName      string
		expectedError error
		cliError      error
	}{
		// Test that the create function works with no error
		{"1", nil, nil},
		// Test that unexpected generic error is handled as an error
		{"2", errors.New("fake error"), errors.New("fake error")},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			var drv Kind = Kind{}
			var cli FakeCli = FakeCli{
				CliError: tt.cliError,
			}
			var configuration config.ParticleConfiguration

			err = drv.Destroy(configuration, &cli)
			t.Log(fmt.Sprintf("error: %s", err))

			assert.Equal(t, err, tt.expectedError)
		})
	}
}
