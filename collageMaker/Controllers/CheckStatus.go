package Controllers

import (
	"collageMaker/Global"
	"github.com/gofiber/fiber/v2"
)

func Status(c *fiber.Ctx)error{
	folderId := c.Params("folderId")
	_, prs := Global.TaskBeingDone[folderId]
	if prs==true{
		return c.Status(fiber.StatusProcessing).JSON(fiber.Map{"err":false,"task":"in-progress"})
	}
	_,prs = Global.TaskCompleted[folderId]
	if prs==true {
		link := "http://localhost:3000/api/download/"+folderId
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{"err":false,"task":"completed","downloadLink":link})
	}
	_,prs = Global.TaskError[folderId]
	if prs==true {
		delete(Global.TaskError,folderId)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"err":false,"task":"error"})
	}
	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err":true})
}