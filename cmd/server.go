// `toggle server` command for launching the API server
package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func NewServerCommand() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "Run the API server",
		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("Running server")
			return nil
		},
	}
}
