package renderer

import (
	"fmt"
	"image"

	"emu/constants"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var imgFunc func() image.Image
var keypressFunc func(key constants.Key)
var keyReleaseFunc func(key constants.Key)

type Game struct {
	image image.Image
	keys  []ebiten.Key
}

func (g *Game) Update() error {
	g.image = imgFunc()
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.image == nil {
		return
	}

	for _, p := range allKeys {
		var pressed bool

		for _, key := range g.keys {
			if key == keyToEbiten(p) {
				keypressFunc(p)
				pressed = true
			}
		}

		if !pressed {
			keyReleaseFunc(p)
		}
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(10, 10)

	screen.DrawImage(ebiten.NewImageFromImage(g.image), op)

	// display fps
	msg := fmt.Sprintf(`TPS: %0.2f | FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func (g *Game) Start(width, height int) error {
	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Chip8")
	ebiten.SetMaxTPS(1000)
	if err := ebiten.RunGame(&Game{}); err != nil {
		return err
	}
	return nil
}

func SetImgFunc(fn func() image.Image) {
	imgFunc = fn
}

func SetKeypressFunc(fn func(key constants.Key)) {
	keypressFunc = fn
}

func SetKeyReleaseFunc(fn func(key constants.Key)) {
	keyReleaseFunc = fn
}

func New() *Game {
	return &Game{}
}
