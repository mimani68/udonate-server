package middleware

import (
	"udonate/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func GeneralMiddleWare(app *fiber.App, configuration config.IConfig) {
	app.Use(recover.New())
	corsConfig := cors.Config{
		// AllowOrigins: "https://gofiber.io, https://gofiber.net",
		// AllowHeaders: "Origin, Content-Type, Accept",
	}
	app.Use(cors.New(corsConfig))
	app.Use(etag.New())
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${latency} ${method} ${path}\n",
		TimeFormat: "2006-02-31T15:04:05",
		TimeZone:   configuration.Get("TZ"),
	}))
}
