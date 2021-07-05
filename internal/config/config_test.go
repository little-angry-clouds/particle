package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateConfiguration(t *testing.T) {
	var test = []struct {
		testName string
		path     string
	}{
		{
			"1",
			"particle_1",
		},
		{
			"2",
			"particle_2",
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
			fullPathFile := fmt.Sprintf("%s/particle.yml", path)

			// Create the test directory
			_ = os.MkdirAll(path, 0755)

			t.Log(fmt.Sprintf("full path: %s", fullPathFile))
			err = CreateConfiguration(path, config)

			// Check there's no error when executing the function
			assert.Nil(t, err)

			// Check the configuration file exists and is not a directory
			exists, err := os.Stat(fullPathFile)
			assert.Nil(t, err)
			assert.Equal(t, false, exists.IsDir())
		})
	}
}
