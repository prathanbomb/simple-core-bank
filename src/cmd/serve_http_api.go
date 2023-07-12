package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/oatsaysai/simple-core-bank/src/app"
	"github.com/oatsaysai/simple-core-bank/src/http_api"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveBackOfficeAPICmd)
}

var serveBackOfficeAPICmd = &cobra.Command{
	Use:   "serve-http-api",
	Short: "Start HTTP API server",
	RunE: func(cmd *cobra.Command, args []string) error {

		logger, err := getLogger()
		if err != nil {
			return err
		}

		app, err := app.New(logger)
		if err != nil {
			return err
		}

		fiberApp := fiber.New()
		httpAPI, err := http_api.New(fiberApp, app)
		if err != nil {
			return err
		}

		return fiberApp.Listen(fmt.Sprintf(":%d", httpAPI.Config.Port))
	},
}
