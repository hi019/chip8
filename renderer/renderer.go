package renderer

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var imgFunc func() image.Image

type Game struct {
	image image.Image
}

func (g *Game) Update() error {
	g.image = imgFunc()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.image == nil {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)

	screen.DrawImage(ebiten.NewImageFromImage(g.image), op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *Game) Start(width, height int) error {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Chip8")
	if err := ebiten.RunGame(&Game{}); err != nil {
		return err
	}
	return nil
}

func SetImgFunc(fn func() image.Image) {
	imgFunc = fn
}

func New() *Game {
	return &Game{}
}
