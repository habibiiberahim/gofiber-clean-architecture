package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func FiberMiddleware(a *fiber.App) {
	a.Use(
		cors.New(),
		compress.New(),
		etag.New(),
		favicon.New(),
		recover.New(),
	)
}
