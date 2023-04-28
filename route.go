package main

import (
	"fmt"
	"net/http"
	"time"

	"api.default.indicoinnovation.pt/entity"
	"api.default.indicoinnovation.pt/handler/health"
	"api.default.indicoinnovation.pt/middleware"
	"api.default.indicoinnovation.pt/pkg/app"
	"api.default.indicoinnovation.pt/pkg/constants"
	"api.default.indicoinnovation.pt/pkg/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func route() *fiber.App {
	allowedOrigins := constants.AllowedOrigins
	if app.Inst.Config.Environment != "production" {
		allowedOrigins += fmt.Sprintf(", %s", constants.AllowedStageOrigins)
	}

	app.Inst.Server.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,OPTIONS",
		AllowOrigins: allowedOrigins,
		AllowHeaders: "X-Session-Id, Authorization, Content-Type, Accept, Origin",
	}))
	app.Inst.Server.Use(recover.New())
	app.Inst.Server.Use(logger.New())
	app.Inst.Server.Use(middleware.SecurityHeaders())
	app.Inst.Server.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression,
	}))

	apiGroup := app.Inst.Server.Group("/api")
	// apiGroup.Use(middleware.Authorize())

	apiGroup.Use(limiter.New(limiter.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.IP() == "127.0.0.1"
		},
		Max:        constants.MaxResquestLimit,
		Expiration: 1 * time.Minute,
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.Get("x-forwarded-for")
		},
		LimitReached: func(c *fiber.Ctx) error {
			return helpers.CreateResponse(c, &entity.ErrorResponse{
				Message:     "Calls Limit Reached",
				Description: "Rate Limit reached",
				StatusCode:  http.StatusTooManyRequests,
			}, http.StatusTooManyRequests)
		},
	}))

	apiGroup.Get("/health", health.Handle().Check)

	return app.Inst.Server
}
