// `toggle client` command for calling the Toggle API server from the command line
package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func NewClientCommand() *cli.Command {
	return &cli.Command{
		Name:    "client",
		Aliases: []string{"c"},
		Usage:   "Client for interacting with the API server",
		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("Running client")
			return nil
		},
	}
}
