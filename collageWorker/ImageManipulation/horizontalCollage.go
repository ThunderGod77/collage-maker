package ImageManipulation

import (
	"collageWorker/PublishResult"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

func HorizontalCollage(image1 image.Image, image2 image.Image, image3 image.Image, borderWidth int, borderColor color.RGBA, folderName string) {
	//starting point of image 1
	sp1 := image.Point{
		X: borderWidth,
		Y: borderWidth,
	}
	//rectangle to enclose first image
	r1 := image.Rectangle{
		Min: sp1,
		Max: sp1.Add(image1.Bounds().Size()),
	}
	//starting point of image 2
	sp2 := image.Point{
		X: sp1.X + image1.Bounds().Size().X + borderWidth,
		Y: borderWidth,
	}
	// rectangle to enclose second image
	r2 := image.Rectangle{
		Min: sp2,
		Max: sp2.Add(image2.Bounds().Size()),
	}
	//starting point of image 3
	sp3 := image.Point{
		X: sp2.X + image2.Bounds().Size().X + borderWidth,
		Y: borderWidth,
	}
	// rectangle to enclose third image
	r3 := image.Rectangle{
		Min: sp3,
		Max: sp3.Add(image3.Bounds().Size()),
	}
	// rectangle to enclose all images
	r := image.Rectangle{
		Min: image.Point{},
		Max: r3.Max.Add(image.Point{
			X: borderWidth,
			Y: borderWidth,
		}),
	}

	rgba := image.NewRGBA(r)
	draw.Draw(rgba, rgba.Bounds(), &image.Uniform{C: borderColor}, image.Point{}, draw.Src)
	draw.Draw(rgba, r1, image1, image.Point{}, draw.Src)
	draw.Draw(rgba, r2, image2, image.Point{}, draw.Src)
	draw.Draw(rgba, r3, image3, image.Point{}, draw.Src)
	out, err := os.Create("./../" + folderName + "/output.jpg")
	if err != nil {
		fmt.Println(err)
		PublishResult.OnError("Error creating collage!",folderName)
		return
	}

	var opt jpeg.Options
	opt.Quality = 80

	err = jpeg.Encode(out, rgba, &opt)
	if err != nil {
		log.Println(err)
		PublishResult.OnError("Error encoding image!",folderName)
	}
	PublishResult.OnCompletion(folderName)
}
