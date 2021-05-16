package Controllers

import (
	"collageMaker/Global"
	"github.com/gofiber/fiber/v2"
	"os"
)

func DownloadLink(c *fiber.Ctx) error {
	folderId := c.Params("folderId")

	_, prs := Global.TaskCompleted[folderId]
	if prs == true {
		_, err := os.Stat("./../" + folderId + "/output.jpg")
		if os.IsNotExist(err) {
			return c.Status(fiber.StatusGone).JSON(fiber.Map{"msg":"Download link expired!"})
		}
		c.Set("Content-Disposition", "attachment; filename=collage.jpg")
		c.Set("Content-Type", "application/octet-stream")
		return c.SendFile("./../" + folderId + "/output.jpg")
	} else {
		return c.Status(fiber.StatusGone).JSON(fiber.Map{"msg":"Download link expired!"})
	}

}
