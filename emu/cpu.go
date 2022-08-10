package emu

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"

	"emu/constants"
)

type CPU struct {
	Registers [16]byte
	Memory    [4096]byte
	Key       [16]byte
	Video     [constants.VideoHeight * constants.VideoWidth]uint32
	Opcode    uint16
	// the program counter, which tells us what instruction the cpu is currently executing
	PC uint16
	// stack pointer
	SP    uint16
	Stack [16]uint16
	// delay timer
	DT byte
	// sound timer
	ST byte
	// index
	I      uint16
	Keypad [16]uint8
}

func (c *CPU) Init() {
	c.PC = 0x200

	// init fonts, starting at 0x50
	for i, b := range fontset {
		c.Memory[FontsetStartAddress+i] = b
	}
}

func (c *CPU) LoadROM(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}

	for i := 0; true; i++ {
		// rom starts at 0x200
		_, err := f.Read(c.Memory[0x200:])
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}

	return nil
}

func (c *CPU) Random() uint8 {
	return uint8(rand.Intn(255))
}

func (c *CPU) Cycle() {
	//if int(c.PC) == len(c.Memory) {
	//	c.PC = 0x200
	//}

	// fetch
	opcode := (uint16(c.Memory[c.PC]) << 8) | uint16(c.Memory[c.PC+1])
	// increment pc
	c.PC += 2

	inHex := fmt.Sprintf("%x", opcode)
	log.Println(inHex)

	// decode and execute
	c.Opcode = opcode
	MainMap[opcode&0xF000>>12](opcode)(opcode, c)

	// decrease delay timer if it's set
	if c.DT > 0 {
		c.DT--
	}
	// decrease sound timer if it's set
	if c.ST > 0 {
		c.ST--
	}
}
