package Routes

import (
	"github.com/gofiber/fiber/v2"

	"collageMaker/Controllers"
)

func ApiRoute(route fiber.Router) {
	route.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"ping": "pong"})
	})
	route.Get("/status/:folderId",Controllers.Status)
	route.Get("/download/:folderId",Controllers.DownloadLink)
	route.Post("/upload", Controllers.ImageUpload)

}
