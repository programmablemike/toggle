// `toggle client` command for calling the Toggle API server from the command line
package cmd

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	togglev1 "github.com/programmablemike/toggle/gen/go/toggle/v1"
	"github.com/programmablemike/toggle/gen/go/toggle/v1/togglev1connect"

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

			client := togglev1connect.NewToggleServiceClient(
				http.DefaultClient,
				"http://localhost:8081",
			)

			res, err := client.CreateScope(
				context.Background(),
				connect.NewRequest(&togglev1.CreateScopeRequest{}),
			)
			if err != nil {
				log.Error().Err(err)
				return err
			}
			log.Info(res)

			return nil
		},
	}
}
