package main

import (
	"crawler-lostark/pkg/cmd"
	"crawler-lostark/pkg/config"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var (
	version = "development"
)

func main() {
	app := cli.NewApp()
	app.Name = "crawler-lostark"
	app.Version = version
	app.Flags = config.InitConfiguration()
	app.Commands = cmd.Cmd()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
