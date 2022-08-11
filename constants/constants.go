package constants

const (
	VideoHeight = 32
	VideoWidth  = 64
)

type Key int

const (
	Key1 Key = iota
	Key2 Key = iota
	Key3 Key = iota
	Key4 Key = iota
	KeyQ Key = iota
	KeyW Key = iota
	KeyE Key = iota
	KeyR Key = iota
	KeyA Key = iota
	KeyS Key = iota
	KeyD Key = iota
	KeyF Key = iota
	KeyZ Key = iota
	KeyX Key = iota
	KeyC Key = iota
	KeyV Key = iota
)

func (k Key) ToChip8() uint8 {
	switch k {
	case Key1:
		return 0x1
	case Key2:
		return 0x2
	case Key3:
		return 0x3
	case Key4:
		return 0xC
	case KeyQ:
		return 0x4
	case KeyW:
		return 0x5
	case KeyE:
		return 0x6
	case KeyR:
		return 0xD
	case KeyA:
		return 0x7
	case KeyS:
		return 0x8
	case KeyD:
		return 0x9
	case KeyF:
		return 0xE
	case KeyZ:
		return 0xA
	case KeyX:
		return 0x0
	case KeyC:
		return 0xB
	case KeyV:
		return 0xF
	default:
		panic("invalid key")
	}
}
