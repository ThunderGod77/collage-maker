package ImageManipulation

import (
	"collageWorker/Global"
	"collageWorker/PublishResult"
	"image"
	"image/color"
	"log"
	"os"
	"strconv"
)

func ImagePreProcessing(imgInfo Global.ImagesInfo) {
	reader1, err := os.Open("./../" + imgInfo.FolderId + "/" + imgInfo.Image1)
	if err != nil {
		log.Println(err)
		PublishResult.OnError("Error downloading image!",imgInfo.FolderId)
		return
	}
	defer reader1.Close()
	reader2, err := os.Open("./../" + imgInfo.FolderId + "/" + imgInfo.Image2)
	if err != nil {
		log.Println(err)
		PublishResult.OnError("Error downloading image!",imgInfo.FolderId)
		return
	}
	defer reader2.Close()
	reader3, err := os.Open("./../" + imgInfo.FolderId + "/" + imgInfo.Image3)
	if err != nil {
		log.Println(err)
		PublishResult.OnError("Error downloading image!",imgInfo.FolderId)
		return
	}
	defer reader3.Close()
	a, _, err := image.Decode(reader1)
	if err != nil {
		log.Println(err)
		PublishResult.OnError("Error decoding image!",imgInfo.FolderId)
		return
	}
	b, _, err := image.Decode(reader2)
	if err != nil {
		log.Println(err)
		PublishResult.OnError("Error decoding image!",imgInfo.FolderId)
		return
	}
	c, _, err := image.Decode(reader3)
	if err != nil {
		log.Println(err)
		PublishResult.OnError("Error decoding image!",imgInfo.FolderId)
		return
	}
	a, b, c = ResizeImage(a, b, c)
	klr := imgInfo.Color
	red, _ := strconv.Atoi(klr[0:3])
	green, _ := strconv.Atoi(klr[3:6])
	blue, _ := strconv.Atoi(klr[6:9])
	alpha, _ := strconv.Atoi(klr[9:12])
	clr := color.RGBA{
		R: uint8(red),
		G: uint8(green),
		B: uint8(blue),
		A: uint8(alpha),
	}
	bw, _ := strconv.Atoi(imgInfo.BorderWidth)
	if imgInfo.Instruction == "horizontal" {
		go HorizontalCollage(a, b, c, bw, clr, imgInfo.FolderId)
	}else if imgInfo.Instruction == "vertical" {
		go VerticalCollage(a,b,c,bw,clr,imgInfo.FolderId)
	}


}
