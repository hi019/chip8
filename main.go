package main

import (
	"embed"
	"image"
	"log"

	"emu/constants"
	"emu/emu"
	"emu/renderer"
)

//go:embed assets/*
var assets embed.FS

func main() {
	c := emu.CPU{}
	c.Init()

	f, err := assets.Open("assets/roms/PONG2")
	if err != nil {
		log.Fatalln(err)
	}
	if err := c.LoadROM(f); err != nil {
		log.Fatalln(err)
	}

	ren := renderer.New()
	renderer.SetKeypressFunc(func(key constants.Key) {
		c.Keypad[key.ToChip8()] = 1
	})
	renderer.SetKeyReleaseFunc(func(key constants.Key) {
		c.Keypad[key.ToChip8()] = 0
	})
	renderer.SetImgFunc(func() image.Image {
		c.Cycle()

		img := image.NewGray(image.Rect(0, 0, constants.VideoWidth, constants.VideoHeight))

		for i, _ := range c.Video {
			img.Pix[i] = c.Video[i]
		}
		return img
	})

	if err := ren.Start(640, 320); err != nil {
		log.Fatalln(err)
	}

}
