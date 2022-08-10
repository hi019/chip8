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

	if err := c.LoadROM("./test_opcode.ch8"); err != nil {
		log.Fatalln(err)
	}

	ren := renderer.New()
	renderer.SetImgFunc(func() image.Image {
		c.Cycle()

		img := image.NewGray(image.Rect(0, 0, constants.VideoWidth, constants.VideoHeight))

		for i, u := range c.Video {
			img.Pix[i] = uint8(u)
		}
		return img
	})

	if err := ren.Start(640, 480); err != nil {
		log.Fatalln(err)
	}

}
