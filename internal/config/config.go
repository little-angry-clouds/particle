package config

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Key string

type ParticleConfiguration struct {
	Driver      Driver      `yaml:"driver"`
	Provisioner Provisioner `yaml:"provisioner"`
	Lint        string      `yaml:"lint"`
}

type Driver struct {
	Name              string `yaml:"name"`
	KubernetesVersion Key    `yaml:"kubernetes_version"`
}

type Provisioner struct {
	Name string `yaml:"name"`
}

func CreateConfiguration(path string, scenario string, configuration ParticleConfiguration) error {
	var configDirPath string = path + "/particle/" + scenario + "/"
	var configFilePath string = configDirPath + "particle.yml"
	var err error

	// Check if the directory exists
	_, err = os.Stat(configDirPath)
	if !os.IsNotExist(err) {
		return errors.New("particle already initialiazed")
	}

	// Create directory
	err = os.MkdirAll(configDirPath, 0755)
	if err != nil {
		return err
	}

	// Create the configuration file
	conf, err := yaml.Marshal(configuration)
	if err != nil {
		return err
	}

	f, err := os.Create(configFilePath)
	if err != nil {
		return err
	}

	defer f.Close()

	_, err = f.Write(conf)
	if err != nil {
		return err
	}

	return nil
}

func ReadConfiguration(scenario string) (ParticleConfiguration, error) {
	var configuration ParticleConfiguration = ParticleConfiguration{}
	var configDirPath string = "particle/" + scenario + "/"
	var configFilePath string = configDirPath + "particle.yml"
	var err error

	// Check if the directory exists
	_, err = os.Stat(configDirPath)
	if os.IsNotExist(err) {
		return configuration, errors.New("particle is not initialiazed")
	}

	// Read the configuration file
	configBinary, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		return configuration, err
	}

	// Create the configuration file
	err = yaml.Unmarshal(configBinary, &configuration)
	if err != nil {
		return configuration, err
	}

	return configuration, nil
}
