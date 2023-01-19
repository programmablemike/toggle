package main

import (
	"os"

	"github.com/rs/zerolog/log"

	"github.com/programmablemike/toggle/cmd"
)

func main() {
	app := cmd.NewApp()

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err)
	}
}
