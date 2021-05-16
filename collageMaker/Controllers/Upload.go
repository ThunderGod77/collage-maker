package Controllers

import (
	"bytes"
	"collageMaker/Global"
	"collageMaker/RabbitMq"

	"encoding/gob"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"log"
	"mime/multipart"
	"os"
)

type fdError struct {
	err bool
	msg string
}
var h string
func fileDownload(files []*multipart.FileHeader, c *fiber.Ctx, folderName string) fdError {

	for _, file := range files {
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

		if !(file.Header["Content-Type"][0] == "image/png" || file.Header["Content-Type"][0] == "image/jpg" || file.Header["Content-Type"][0] == "image/jpeg") {
			return fdError{
				err: true,
				msg: "Wrong File Type!",
			}
		}

		err := c.SaveFile(file, fmt.Sprintf("./../%s/%s", folderName, file.Filename))
		if err != nil {
			log.Println(err)
			return fdError{
				err: true,
				msg: "Error saving file!",
			}
		}

	}
	return fdError{
		err: false,
		msg: "Saved successfully",
	}
}

func ImageUpload(c *fiber.Ctx) error {
	if form, err := c.MultipartForm(); err == nil {
		// => *multipart.Form

		if token := form.Value["token"]; len(token) > 0 {
			// Get key value:
			fmt.Println(token[0])
		}
		instruction := form.Value["instruction"]
		if len(instruction) < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": true, "msg": "No instructions provided!"})
		}
		instr := instruction[0]
		color := form.Value["color"]
		if len(color) < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": true, "msg": "Invalid Color!"})
		}
		clr := color[0]
		borderWidth := form.Value["borderWidth"]
		if len(borderWidth) < 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": true, "msg": "Invalid border width!"})
		}
		bwt := borderWidth[0]

		// Get all files from "documents" key:
		files := form.File["documents"]
		// => []*multipart.FileHeader
		if len(files) != 3 {
			msg := "Incorrect number of files attached"
			log.Println(msg)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"err": true, "msg": msg})
		}
		// Loop through files:
		h = utils.UUIDv4()
		err := os.Mkdir("./../"+h, 0755)
		if err != nil {
			log.Fatal(err)
		}
		fde := fileDownload(files, c, h)
		if fde.err == true {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": true, "msg": fde.msg})
		}
		imgInfo := Global.ImagesInfo{
			Image1:      files[0].Filename,
			Image2:      files[1].Filename,
			Image3:      files[2].Filename,
			FolderId:    h,
			Instruction: instr,
			Color:       clr,
			BorderWidth: bwt,
		}

		var mqCarrier bytes.Buffer
		enc := gob.NewEncoder(&mqCarrier)
		err = enc.Encode(imgInfo)
		if err != nil {
			log.Println(err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"err": true, "msg": "Error encoding data!"})
		}
		RabbitMq.Publish(mqCarrier.Bytes(),h)
		Global.TaskBeingDone[h] = true
		RabbitMq.PublishDelete([]byte(h))

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"err": false,"id":h})
}
