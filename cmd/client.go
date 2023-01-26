// `toggle client` command for calling the Toggle API server from the command line
package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	togglev1 "github.com/programmablemike/toggle/gen/go/toggle/v1"
	"github.com/programmablemike/toggle/gen/go/toggle/v1/togglev1connect"

	"github.com/rs/zerolog/log"
	uuid "github.com/satori/go.uuid"
	"github.com/urfave/cli/v2"
)

func CreateToggleClient(cCtx *cli.Context) togglev1connect.ToggleServiceClient {
	return togglev1connect.NewToggleServiceClient(
		http.DefaultClient,
		cCtx.String("protocol")+cCtx.String("address"),
	)
}

func NewClientCommand() *cli.Command {
	return &cli.Command{
		Name:    "client",
		Aliases: []string{"c"},
		Usage:   "Client for interacting with the API server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "protocol",
				Value: "http://",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "create-scope-set",
				Usage: "Create a new scope set",
				Action: func(cCtx *cli.Context) error {
					log.Info().Msg("Calling `create scope set`")

					client := CreateToggleClient(cCtx)
					res, err := client.CreateScopeSet(
						context.Background(), // TODO(mlee): Add a timeout to the context
						connect.NewRequest(&togglev1.CreateScopeSetRequest{
							Info: &togglev1.MessageInfo{
								Id: fmt.Sprintf("%s", uuid.NewV4()),
							},
						}),
					)

					if err != nil {
						log.Error().Err(err)
						return err
					}
					log.Info().Msgf("Received response: %v", res)
					return nil
				},
			},
			{
				Name:  "create-scope",
				Usage: "Create a new scope",
				Action: func(cCtx *cli.Context) error {
					log.Info().Msg("Calling `create scope`")

					client := CreateToggleClient(cCtx)
					res, err := client.CreateScope(
						context.Background(), // TODO(mlee): Add a timeout to the context
						connect.NewRequest(&togglev1.CreateScopeRequest{
							Info: &togglev1.MessageInfo{
								Id: fmt.Sprintf("%s", uuid.NewV4()),
							},
						}),
					)
					if err != nil {
						log.Error().Err(err)
						return err
					}
					log.Info().Msgf("Received response: %v", res)
					return nil
				},
			},
		},
	}
}
