package verifier

import (
	"github.com/apex/log"
	"github.com/little-angry-clouds/particle/internal/cmd"
	"github.com/little-angry-clouds/particle/internal/config"
)

// Command is an arbitrary implementation of a verifier. It really is just a wrapper
// to execute arbitrary commands.
type Command struct {
	Logger *log.Entry
}

// Verify executes all the commands from its section like 'bash -c "$commands"'.
func (c *Command) Verify(configuration config.ParticleConfiguration, cmd cmd.Cmd) error {
	var logger *log.Entry = c.Logger
	var err error

	args := []string{"bash", "-c", configuration.Verifier}

	err = cmd.Initialize(logger, args)
	if err != nil {
		return err
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return err
}
