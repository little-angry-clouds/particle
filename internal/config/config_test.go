package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfiguration(t *testing.T) { // nolint: funlen
	var test = []struct {
		testName string
		path     string
		scenario string
	}{
		{
			"1",
			"particle_1",
			"default",
		},
		{
			"2",
			"particle_2",
			"whatever",
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.testName, func(t *testing.T) {
			var err error
			var config ParticleConfiguration = ParticleConfiguration{}

			dir, err := ioutil.TempDir("/tmp/", "particle")
			assert.Nil(t, err)
			defer os.Remove(dir)

			path := fmt.Sprintf("%s/%s", dir, tt.path)
			fullPathDir := fmt.Sprintf("%s/particle/%s/", path, tt.scenario)
			fullPathFile := fmt.Sprintf("%s/particle/%s/particle.yml", path, tt.scenario)

			t.Log(fmt.Sprintf("full path: %s", fullPathDir))
			err = CreateConfiguration(path, tt.scenario, config)

			// Check there's no error when executing the function
			assert.Nil(t, err)

			// Check the path directory exists
			exists, err := os.Stat(fullPathDir)
			assert.Nil(t, err)
			assert.Equal(t, true, exists.IsDir())

			// Check the configuration file exists and is not a directory
			exists, err = os.Stat(fullPathFile)
			assert.Nil(t, err)
			assert.Equal(t, false, exists.IsDir())
		})
	}
}
