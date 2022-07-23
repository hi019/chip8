package main

// Registers contains 8 values, each able to hold up to 8 bits (1 byte)
type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	// f is a "flags" register: the lower 4 bits are always 0, and the cpu automatically writes to the upper four
	// when specific things happen
	f *FlagRegister
	h uint8
	l uint8
}

// VIRTUAL REGISTRIES
// ------------------
// while each register contains 8 bits (an 8 bit cpu), games have the ability to read and write 16 bits at once
// into these "virtual" registers, which are actually just two registers combined. there are 4 of these:
//  - af (registry a and f)
//	- bc
//	- de
//	- hl
//
// to read a virtual registry, we add it's two byte components together. eg 0b11001100 + 0b10110010 = 0b110011000b10110010
// how this is accomplished:
// 1. convert both registries (components) to u16. this effectively adds a byte (or however many bits is required to make 16 bits)
// of 0s to the left of each component value.
// 2. shift the first component 8 bytes to the left
// 3. subtract both components

func (r *Registers) GetBC() uint16 {
	return (uint16(r.b) << 8) | uint16(r.c)
}
