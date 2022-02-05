package config

import "github.com/urfave/cli/v2"

var (
	// Port - application http port
	Port int

	// Database - database connection string
	Database       string
	DatabaseEngine string

	// Debug - for enable debug mode
	Debug bool

	// Discod webhook
	DiscordWebhook string
)

func InitConfiguration() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:        "port",
			Aliases:     []string{"p"},
			EnvVars:     []string{"PORT"},
			Value:       8080,
			Usage:       "Server Port",
			Destination: &Port,
		},
		&cli.StringFlag{
			Name:        "database",
			Aliases:     []string{"d"},
			EnvVars:     []string{"DATABASE"},
			Destination: &Database,
		},
		&cli.StringFlag{
			Name:        "database_engine",
			Aliases:     []string{"de"},
			EnvVars:     []string{"DATABASE_ENGINE"},
			Destination: &DatabaseEngine,
		},
		&cli.StringFlag{
			Name:        "discord_webhook",
			Aliases:     []string{"dw"},
			EnvVars:     []string{"DISCORD_WEBHOOK"},
			Destination: &DiscordWebhook,
		},
		&cli.BoolFlag{
			Name:        "debug",
			EnvVars:     []string{"DEBUG"},
			Destination: &Debug,
		},
	}
}
