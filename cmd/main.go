package cmd

import (
	"github.com/urfave/cli/v2"
)

func NewApp() *cli.App {
	return &cli.App{
		Name:  "toggle",
		Usage: "feature flagging API and client",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "address",
				Aliases: []string{"a"},
				Value:   "127.0.0.1:8080",
			},
		},
		Commands: []*cli.Command{
			NewServerCommand(),
			NewClientCommand(),
		},
	}
}
