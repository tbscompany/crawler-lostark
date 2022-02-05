package cmd

import (
	r "crawler-lostark/pkg/runner"

	"github.com/urfave/cli/v2"
)

func runner() *cli.Command {
	return &cli.Command{
		Name:   "runner",
		Action: r.StartRunner,
	}
}
