package verifier

import (
	"errors"
	"fmt"
	"testing"

	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/config"
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

func TestVerify(t *testing.T) {
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
			var vrf Helm = Helm{}
			var cli FakeCli = FakeCli{
				CliError: tt.cliError,
			}
			var configuration config.ParticleConfiguration

			err = vrf.Verify(configuration, &cli)
			t.Log(fmt.Sprintf("error: %s", err))

			assert.Equal(t, err, tt.expectedError)
		})
	}
}
