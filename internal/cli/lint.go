package cli

import (
	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

func Lint(scenario string, configuration config.ParticleConfiguration, logger *log.Entry) error {
	var err error
	var cli cmd.CLI = cmd.CLI{Binary: "bash"}

	cmdArgs := []string{"bash", "-c", configuration.Linter}

	err = cli.Initialize(logger, cmdArgs)
	if err != nil {
		return err
	}

	err = cli.Run()
	if err != nil {
		return err
	}

	return err
}
