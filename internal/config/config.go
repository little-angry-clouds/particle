package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type ParticleConfiguration struct {
	Driver      Driver      `yaml:"driver" validate:"required"`
	Provisioner Provisioner `yaml:"provisioner" validate:"required"`
	Linter      string      `yaml:"lint" validate:"required"`
	Verifier    string      `yaml:"verifier" validate:"required"`
	Dependency  Dependency  `yaml:"dependency" validate:"required,eq=helm"`
	Prepare     []Prepare   `yaml:"prepare,omitempty"`
}

type Driver struct {
	Name              string                 `yaml:"name" validate:"required,eq=kind|eq=minikube"`
	KubernetesVersion string                 `yaml:"kubernetes-version,omitempty"`
	Values            map[string]interface{} `yaml:"values,omitempty"`
}

type Provisioner struct {
	Name   string                 `yaml:"name" validate:"eq=helm"`
	Values map[string]interface{} `yaml:"values,omitempty"`
}

type Dependency struct {
	Name   string         `yaml:"name" validate:"eq=helm"`
	Charts []Dependencies `yaml:"charts,omitempty"`
}

type Dependencies struct {
	RepositoryName string `yaml:"repository-name,omitempty"`
	RepositoryURL  string `yaml:"repository-url,omitempty"`
}

type Prepare struct {
	Name    string                 `yaml:"name"`
	Version string                 `yaml:"version,omitempty"`
	Values  map[string]interface{} `yaml:"values,omitempty"`
}

func CreateConfiguration(configDirPath string, configuration ParticleConfiguration) error {
	var configFilePath string = configDirPath + "/particle.yml"
	var err error

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
		return configuration, &particleNotInitialized{}
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
