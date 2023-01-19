package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:  "toggle",
		Usage: "feature flagging API and client",
		Commands: []*cli.Command{
			NewServerCommand(),
			NewClientCommand(),
		},
	}
}
