package http_api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/oatsaysai/simple-core-bank/src/app"
	"github.com/oatsaysai/simple-core-bank/src/http_api/routes"
	log "github.com/oatsaysai/simple-core-bank/src/logger"
)

type HttpAPI struct {
	Config *Config
	App    *app.App
}

func New(fiberApp *fiber.App, app *app.App) (httpAPI *HttpAPI, err error) {
	httpAPI = &HttpAPI{
		App: app,
	}
	httpAPI.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	fiberApp.Use(cors.New())
	fiberApp.Use(httpAPI.loggingMiddleware())

	api := fiberApp.Group("/api")
	routes.DefaultRouter(api)

	return httpAPI, nil
}

func (api *HttpAPI) loggingMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		startTime := time.Now()
		logger := api.App.Logger.WithFields(log.Fields{
			"package":   "http_api",
			"remote_ip": c.Context().RemoteIP().String(),
		})

		c.Next()

		duration := time.Since(startTime)
		statusCode := c.Response().StatusCode()
		logger = logger.WithFields(log.Fields{
			"duration":    duration.String(),
			"status_code": statusCode,
		})
		logger.Infof("%s %s", c.Method(), c.OriginalURL())

		return nil
	}
}
