package cmd

import (
	d "crawler-lostark/pkg/database"
	"crawler-lostark/pkg/logger"
	"crawler-lostark/pkg/models"

	"github.com/urfave/cli/v2"
)

func database() *cli.Command {
	return &cli.Command{
		Name: "database",
		Subcommands: []*cli.Command{
			{
				Name:   "migrate",
				Action: migrate,
			},
		},
	}
}

func migrate(ctx *cli.Context) error {
	// Init logger
	logger.InitLogger(ctx.App.Version)

	// Init database
	d.InitDatabase()

	c := d.GetDbConnection()

	logger.Logger.Info().Msg("Init/Migrate database scheme")
	err := c.DB.AutoMigrate(&models.News{})
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Unable to init/migrate database scheme")
		return nil
	}
	return nil
}
