package http_api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/oatsaysai/simple-core-bank/src/http_api/routes"
)

type HttpAPI struct {
	Config *Config
}

func New(fiberApp *fiber.App) (httpAPI *HttpAPI, err error) {
	httpAPI = &HttpAPI{}
	httpAPI.Config, err = InitConfig()
	if err != nil {
		return nil, err
	}

	api := fiberApp.Group("/api")
	routes.DefaultRouter(api)

	return httpAPI, nil
}
