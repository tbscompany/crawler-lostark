package cmd

import "github.com/urfave/cli/v2"

func Cmd() []*cli.Command {
	var c []*cli.Command
	c = append(c, database())
	c = append(c, runner())
	return c
}
