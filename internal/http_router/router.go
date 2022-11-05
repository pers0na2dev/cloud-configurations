package http_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
)

func Router() *fiber.App {
	app := fiber.New(fiber.Config{
		Prefork: true,
		AppName: "cloud-configuration microservice",
	})

	app.Get("/metrics", monitor.New(monitor.Config{
		Title: "cloud-configuration microservice",
	}))

	v1 := app.Group("/api/v1")
	v1.Post("/getGet", GetAll)
	v1.Post("/get", Get)
	v1.Post("/create", Create)
	v1.Post("/save", Save)
	v1.Post("/delete", Delete)

	return app
}
