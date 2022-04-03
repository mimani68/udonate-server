package middleware

import "github.com/gofiber/fiber/v2"

func Auth(app *fiber.App) {
	app.Use(New())
}

func New() fiber.Handler {
	// Set default config
	// cfg := configDefault(config...)

	// Return new handler
	return func(c *fiber.Ctx) (err error) {
		// Don't execute middleware if Next returns true
		c.Set("USER", "SALAM")
		// if cfg.Next != nil && cfg.Next(c) {
		return c.Next()
		// }
	}
}
