package main

import (
	"collageMaker/RabbitMq"
	"collageMaker/Routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "image/jpeg"
	_ "image/png"
)


func setupRoutes(app *fiber.App) {

	Routes.ApiRoute(app.Group("/api"))

}

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept,Authorization,Content-Disposition",
	}))
	app.Use(logger.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"ping": "pong"})
	})
	app.Get("/po", func(ctx *fiber.Ctx) error {

		ctx.Set("Content-Disposition","attachment; filename=lol.png")
		ctx.Set("Content-Type", "application/octet-stream")
		return ctx.SendFile("./lol.png")

	})
	setupRoutes(app)
	RabbitMq.InitRabbitMq()
	defer RabbitMq.Conn.Close()
	defer RabbitMq.Ch.Close()

	app.Listen(":3000")

}

