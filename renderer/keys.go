package renderer

import (
	"emu/constants"
	"github.com/hajimehoshi/ebiten/v2"
)

var allKeys = []constants.Key{
	constants.Key1,
	constants.Key2,
	constants.Key3,
	constants.Key4,
	constants.KeyQ,
	constants.KeyW,
	constants.KeyE,
	constants.KeyR,
	constants.KeyA,
	constants.KeyS,
	constants.KeyD,
	constants.KeyF,
	constants.KeyZ,
	constants.KeyX,
	constants.KeyC,
	constants.KeyV,
}

func keyToEbiten(k constants.Key) ebiten.Key {
	switch k {
	case constants.Key1:
		return ebiten.Key1
	case constants.Key2:
		return ebiten.Key2
	case constants.Key3:
		return ebiten.Key3
	case constants.Key4:
		return ebiten.Key4
	case constants.KeyQ:
		return ebiten.KeyQ
	case constants.KeyW:
		return ebiten.KeyW
	case constants.KeyE:
		return ebiten.KeyE
	case constants.KeyR:
		return ebiten.KeyR
	case constants.KeyA:
		return ebiten.KeyA
	case constants.KeyS:
		return ebiten.KeyS
	case constants.KeyD:
		return ebiten.KeyD
	case constants.KeyF:
		return ebiten.KeyF
	case constants.KeyZ:
		return ebiten.KeyZ
	case constants.KeyX:
		return ebiten.KeyX
	case constants.KeyC:
		return ebiten.KeyC
	case constants.KeyV:
		return ebiten.KeyV
	default:
		panic("unknown key")
	}
}
