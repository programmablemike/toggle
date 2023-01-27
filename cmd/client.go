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
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "name",
						Value:    "my-scope-set",
						Required: false,
					},
				},
				Action: func(cCtx *cli.Context) error {
					log.Info().Msg("calling `CreateScopeSet`")

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
					log.Info().Msgf("received response: %v", res.Msg)
					return nil
				},
			},
			{
				Name:  "create-scope",
				Usage: "Create a new scope",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "scope-set-id",
						Required: true,
					},
					&cli.BoolFlag{
						Name:     "required",
						Value:    false,
						Required: false,
					},
					&cli.StringFlag{
						Name:     "name",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "description",
						Required: false,
					},
					&cli.StringSliceFlag{
						Name:     "value",
						Required: true,
					},
				},
				Action: func(cCtx *cli.Context) error {
					log.Info().Msg("calling `CreateScope`")

					scope_set_id := cCtx.String("scope-set-id")
					name := cCtx.String("name")
					desc := cCtx.String("description")
					req := cCtx.Bool("required")
					values := cCtx.StringSlice("values")

					client := CreateToggleClient(cCtx)
					res, err := client.CreateScope(
						context.Background(), // TODO(mlee): Add a timeout to the context
						connect.NewRequest(&togglev1.CreateScopeRequest{
							Info: &togglev1.MessageInfo{
								Id: fmt.Sprintf("%s", uuid.NewV4()),
							},
							Value: &togglev1.Scope{
								Info: &togglev1.NameInfo{
									Name:        name,
									Description: &desc,
								},
								ScopeSetId:       scope_set_id,
								Required:         req,
								AcceptableValues: values,
							},
						}),
					)

					if err != nil {
						log.Error().Err(err)
						return err
					}
					log.Info().Msgf("received response: %v", res.Msg)
					return nil

				},
			},
			{
				Name:  "list-scope-sets",
				Usage: "List all available scope sets",
				Action: func(cCtx *cli.Context) error {
					log.Info().Msg("calling `ListScopeSet")

					client := CreateToggleClient(cCtx)
					res, err := client.ListScopeSets(
						context.Background(), //TODO(mlee): Add a timeout to the context
						connect.NewRequest(&togglev1.ListScopeSetsRequest{
							Info: &togglev1.MessageInfo{
								Id: fmt.Sprintf("%s", uuid.NewV4()),
							},
						}),
					)
					if err != nil {
						log.Error().Err(err)
						return err
					}
					log.Info().Msgf("received response: %v", res.Msg)
					return nil
				},
			},
			{
				Name:  "list-scopes",
				Usage: "List all available scopes",
				Action: func(cCtx *cli.Context) error {
					log.Info().Msg("calling `ListScope`")

					client := CreateToggleClient(cCtx)
					res, err := client.ListScopes(
						context.Background(), // TODO(mlee): Add a timeout to the context
						connect.NewRequest(&togglev1.ListScopesRequest{
							Info: &togglev1.MessageInfo{
								Id: fmt.Sprintf("%s", uuid.NewV4()),
							},
						}),
					)
					if err != nil {
						log.Error().Err(err)
						return err
					}
					log.Info().Msgf("received response: %v", res.Msg)
					return nil
				},
			},
		},
	}
}
