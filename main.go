package main

import (
	"image"
	"log"

	"emu/constants"
	"emu/emu"
	"emu/renderer"
)

func main() {
	c := emu.CPU{}
	c.Init()
	//videoScale := 10

	if err := c.LoadROM("./c8_test.c8"); err != nil {
		log.Fatalln(err)
	}

	ren := renderer.New()
	renderer.SetImgFunc(func() image.Image {
		c.Cycle()

		img := image.NewGray(image.Rect(constants.VideoWidth, constants.VideoHeight, 0, 0))

		for i, u := range c.Video {
			img.Pix[i] = uint8(u)
		}
		return img
	})

	if err := ren.Start(1200, 500); err != nil {
		log.Fatalln(err)
	}

}
