// `toggle server` command for launching the API server
package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bufbuild/connect-go"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	togglev1 "github.com/programmablemike/toggle/gen/go/toggle/v1"
	"github.com/programmablemike/toggle/gen/go/toggle/v1/togglev1connect"
	"github.com/programmablemike/toggle/internal/storage"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	grpcreflect "github.com/bufbuild/connect-grpcreflect-go"
)

type ToggleStorage struct {
	toggleset storage.Store[togglev1.ToggleSet]
	toggle    storage.Store[togglev1.Toggle]
	scopeset  storage.Store[togglev1.ScopeSet]
	scope     storage.Store[togglev1.Scope]
}

type ToggleServer struct {
	store *ToggleStorage
}

// CreateScope adds a new scope
func (ts *ToggleServer) CreateScope(ctx context.Context, req *connect.Request[togglev1.CreateScopeRequest]) (*connect.Response[togglev1.CreateScopeResponse], error) {
	log.Info().Msgf("received request: %v", req)
	if req.Msg.Value != nil {
		ts.store.scope.AddRef(req.Msg.Value)
	} else {
		log.Info().Msg("did not receive a Scope in the request")
	}
	res := connect.NewResponse(&togglev1.CreateScopeResponse{
		Info: &togglev1.MessageInfo{Id: fmt.Sprintf("%s", uuid.NewV4())},
	})
	return res, nil
}

// CreateScopeSet adds a new scope set for grouping related scopes
func (ts *ToggleServer) CreateScopeSet(ctx context.Context, req *connect.Request[togglev1.CreateScopeSetRequest]) (*connect.Response[togglev1.CreateScopeSetResponse], error) {
	log.Info().Msgf("received request: %v", req)
	if req.Msg.Value != nil {
		ts.store.scopeset.AddRef(req.Msg.Value)
	} else {
		log.Info().Msg("did not receive a ScopeSet in the request")
	}
	res := connect.NewResponse(&togglev1.CreateScopeSetResponse{
		Info: &togglev1.MessageInfo{Id: fmt.Sprintf("%s", uuid.NewV4())},
	})
	return res, nil
}

func (ts *ToggleServer) ListScopeSets(ctx context.Context, req *connect.Request[togglev1.ListScopeSetsRequest]) (*connect.Response[togglev1.ListScopeSetsResponse], error) {
	log.Info().Msgf("received request: %v", req)
	scopesets := ts.store.scopeset.ListAsRef()
	res := connect.NewResponse(&togglev1.ListScopeSetsResponse{
		Info:   &togglev1.MessageInfo{Id: fmt.Sprintf("%s", uuid.NewV4())},
		Result: scopesets,
	})
	return res, nil
}

func (ts *ToggleServer) ListScopes(ctx context.Context, req *connect.Request[togglev1.ListScopesRequest]) (*connect.Response[togglev1.ListScopesResponse], error) {
	log.Info().Msgf("received request: %v", req)
	scopes := ts.store.scope.ListAsRef()
	res := connect.NewResponse(&togglev1.ListScopesResponse{
		Info:   &togglev1.MessageInfo{Id: fmt.Sprintf("%s", uuid.NewV4())},
		Result: scopes,
	})
	return res, nil
}

func NewToggleStorage() *ToggleStorage {
	return &ToggleStorage{
		toggleset: &storage.DataStore[togglev1.ToggleSet]{},
		toggle:    &storage.DataStore[togglev1.Toggle]{},
		scopeset:  &storage.DataStore[togglev1.ScopeSet]{},
		scope:     &storage.DataStore[togglev1.Scope]{},
	}
}

func NewServerCommand() *cli.Command {
	return &cli.Command{
		Name:    "server",
		Aliases: []string{"s"},
		Usage:   "Run the API server",
		Flags:   []cli.Flag{},
		Action: func(cCtx *cli.Context) error {
			log.Info().Msg("running server 🚀")

			toggler := &ToggleServer{
				store: NewToggleStorage(),
			}
			mux := http.NewServeMux()
			reflector := grpcreflect.NewStaticReflector(
				"toggle.v1.ToggleService",
			)
			path, handler := togglev1connect.NewToggleServiceHandler(toggler)
			mux.Handle(grpcreflect.NewHandlerV1(reflector))
			mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector))
			mux.Handle(path, handler)
			http.ListenAndServe(
				cCtx.String("address"),
				h2c.NewHandler(mux, &http2.Server{}),
			)
			return nil
		},
	}
}
