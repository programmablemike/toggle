// `toggle server` command for launching the API server
package cmd

import (
	"context"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	togglev1 "github.com/programmablemike/toggle/gen/go/toggle/v1"
	"github.com/programmablemike/toggle/gen/go/toggle/v1/togglev1connect"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
)

type ToggleServer struct{}

// CreateScope adds a new scope to partition the toggles
func (ts *ToggleServer) CreateScope(ctx context.Context, req *connect.Request[togglev1.CreateScopeRequest]) (*connect.Response[togglev1.CreateScopeResponse], error) {
	log.Info().Msgf("Request headers: %v", req.Header())
	res := connect.NewResponse(&togglev1.CreateScopeResponse{})
	return res, nil
}

func NewServerCommand() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "Run the API server",
		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("Running server")

			toggler := &ToggleServer{}
			mux := http.NewServeMux()
			reflector := grpcreflect.NewStaticReflector(
				"toggle.v1.ToggleService",
			)
			path, handler := togglev1connect.NewToggleServiceHandler(toggler)
			mux.Handle(grpcreflect.NewHandlerV1(reflector))
			mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
			mux.Handle(path, handler)
			http.ListenAndServe(
				"0.0.0.0:8080", // WARNING(mlee): This is insecure and only suitable for testing
				h2c.NewHandler(mux, &http2.Server{}),
			)
			return nil
		},
	}
}
