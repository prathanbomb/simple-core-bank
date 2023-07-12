package cmd

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

		fiberApp := fiber.New()
		fiberApp.Use(cors.New())

		httpAPI, err := http_api.New(fiberApp)
		if err != nil {
			return err
		}

		return fiberApp.Listen(fmt.Sprintf(":%d", httpAPI.Config.Port))
	},
}
