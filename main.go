package main

import (
	"github.com/edersohe/zflogger"
	"github.com/gofiber/fiber"
)

func main() {
	logger := Logger
	app := fiber.New()

	app.Use(zflogger.Middleware(logger, nil))

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})

	logger.Fatal().Err(app.Listen(3000)).Str("tag", "server").Send()
}
