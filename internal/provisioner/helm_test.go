package provisioner

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

func TestConverge(t *testing.T) {
	var test = []struct {
		testName      string
		expectedError error
		cliError      error
		stdErr        string
	}{
		// Test that the function works with no error
		{"1", nil, nil, ""},
		// Test that unexpected generic error is handled as an error
		{"2",
			// expectedError
			errors.New("fake error"),
			// cliError
			errors.New("fake error"),
			// stdErr
			"",
		},
		// Test that chartCantInstall error is not handled as an error
		{"3",
			// expectedError
			nil,
			// cliError
			nil,
			// sdtErr
			"Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused",
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			// TODO mockear el logger
			var drv Helm = Helm{Logger: helpers.GetLogger(false)}
			var cli FakeCli = FakeCli{
				CliError: tt.cliError,
				Stderr:   tt.stdErr,
			}
			var configuration config.ParticleConfiguration
			configuration.Provisioner.Values = map[string]interface{}{}

			err = drv.Converge(configuration, &cli)
			t.Log(fmt.Sprintf("error: %s", err))

			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestCleanup(t *testing.T) {
	var test = []struct {
		testName      string
		expectedError error
		cliError      error
		stdErr        string
		prepare       []config.Prepare
	}{
		// Test that the function works with no error
		{"1", nil, nil, "", nil},
		// Test that unexpected error is handled as an error without prepare
		{"2",
			// expectedError
			errors.New("fake error"),
			// cliError
			errors.New("fake error"),
			// stdErr
			"",
			// prepare
			nil,
		},
		// Test that unexpected error is handled as an error with prepare
		{"3",
			// expectedError
			errors.New("fake error"),
			// cliError
			errors.New("fake error"),
			// stdErr
			"",
			// prepare
			[]config.Prepare{{Name: "stable/nginx"}},
		},
		// Test that chartNotInstalled is not handled as an error without prepare
		{"3",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"Release not loaded",
			// prepare
			nil,
		},
		// Test that chartNotInstalled is not handled as an error with prepare
		{"4",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"Release not loaded",
			// prepare
			[]config.Prepare{{Name: "stable/nginx"}},
		},
		// Test that chartCanDelete is not handled as an error without prepare
		{"5",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused",
			// prepare
			nil,
		},
		// Test that chartCanDelete is not handled as an error with prepare
		{"6",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused",
			// prepare
			[]config.Prepare{{Name: "stable/nginx"}},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			// TODO mockear el logger
			var drv Helm = Helm{Logger: helpers.GetLogger(false)}
			var cli FakeCli = FakeCli{
				CliError: tt.cliError,
				Stderr:   tt.stdErr,
			}
			var configuration config.ParticleConfiguration
			configuration.Prepare = tt.prepare

			err = drv.Cleanup(configuration, &cli)
			t.Log(fmt.Sprintf("error: %s", err))

			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestDependency(t *testing.T) {
	var test = []struct {
		testName      string
		expectedError error
		cliError      error
		stdErr        string
		dependency    config.Dependency
	}{
		// Test that the function works with no error
		{"1",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"",
			// dependency
			config.Dependency{},
		},
		// Test that the function works with no error
		{"2",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"",
			// dependency
			config.Dependency{
				Name: "helm",
				Charts: []config.Dependencies{{
					RepositoryName: "gitlab",
					RepositoryURL:  "https://charts.gitlab.io",
				}},
			},
		},
		// Test that unexpected error is handled as an error with a dependency
		{"3",
			// expectedError
			errors.New("fake error"),
			// cliError
			errors.New("fake error"),
			// stdErr
			"",
			// dependency
			config.Dependency{
				Name: "helm",
				Charts: []config.Dependencies{{
					RepositoryName: "gitlab",
					RepositoryURL:  "https://charts.gitlab.io",
				}},
			},
		},
		// Test that unexpected error is handled as an error with multiple dependencies
		{"4",
			// expectedError
			errors.New("fake error"),
			// cliError
			errors.New("fake error"),
			// stdErr
			"",
			// dependency
			config.Dependency{
				Name: "helm",
				Charts: []config.Dependencies{{
					RepositoryName: "gitlab",
					RepositoryURL:  "https://charts.gitlab.io",
				}, {
					RepositoryName: "bitname",
					RepositoryURL:  "https://charts.bitnami.com/bitnami",
				}},
			},
		},
		// Test that helmRepoExists is not handled as an error with multiple dependencies
		{"5",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"already exists, please specify a different name",
			// prepare
			config.Dependency{
				Name: "helm",
				Charts: []config.Dependencies{{
					RepositoryName: "gitlab",
					RepositoryURL:  "https://charts.gitlab.io",
				}, {
					RepositoryName: "bitname",
					RepositoryURL:  "https://charts.bitnami.com/bitnami",
				}},
			},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			// TODO mockear el logger
			var drv Helm = Helm{Logger: helpers.GetLogger(false)}
			var cli FakeCli = FakeCli{
				CliError: tt.cliError,
				Stderr:   tt.stdErr,
			}
			var configuration config.ParticleConfiguration
			configuration.Dependency = tt.dependency

			err = drv.Dependency(configuration, &cli)
			t.Log(fmt.Sprintf("error: %s", err))

			assert.Equal(t, tt.expectedError, err)
		})
	}
}

func TestPrepare(t *testing.T) {
	var test = []struct {
		testName      string
		expectedError error
		cliError      error
		stdErr        string
		prepare       []config.Prepare
	}{
		// Test that the function works with no error
		{"1", nil, nil, "", nil},
		// Test that the function works with no error
		{"2",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"",
			// prepare
			[]config.Prepare{{Name: "stable/nginx"}},
		},
		// Test that unexpected generic error is handled as an error
		{"3",
			// expectedError
			errors.New("fake error"),
			// cliError
			errors.New("fake error"),
			// stdErr
			"",
			// prepare
			[]config.Prepare{{Name: "stable/nginx"}},
		},
		// Test that chartCantInstall error is not handled as an error
		{"4",
			// expectedError
			nil,
			// cliError
			nil,
			// stdErr
			"Kubernetes cluster unreachable: Get \"http://localhost:8080/version?timeout=32s\": dial tcp 127.0.0.1:8080: connect: connection refused",
			// prepare
			[]config.Prepare{{Name: "stable/nginx"}},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			// TODO mockear el logger
			var drv Helm = Helm{Logger: helpers.GetLogger(false)}
			var cli FakeCli = FakeCli{
				CliError: tt.cliError,
				Stderr:   tt.stdErr,
			}
			var configuration config.ParticleConfiguration
			configuration.Prepare = tt.prepare

			err = drv.Prepare(configuration, &cli)
			t.Log(fmt.Sprintf("error: %s", err))

			assert.Equal(t, tt.expectedError, err)
		})
	}
}
