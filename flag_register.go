package main

// FlagRegister values are effectively the upper 4 bits of a registry value but represented as boolean's for ergonomics
//     ┌-> Carry
// 	 ┌-+> Subtraction
// 	 | |
//	1111 0000
//	| |
//	└-+> Zero
//    └-> Half Carry
type FlagRegister struct {
	// set to true if the result of the operation is equal to 0
	zero bool
	// set to true if the operation was a subtraction.
	subtract bool
	// set to true if there is an overflow from the lower nibble (a.k.a the lower four bits) to the upper nibble
	// (a.k.a the upper four bits)
	halfCarry bool
	// set to true if the operation resulted in an overflow
	carry bool
}

const (
	ZeroFlagBytePosition      = 7
	SubtractFlagBytePosition  = 6
	HalfCarryFlagBytePosition = 5
	CarryFlagBytePosition     = 4
)

// ToUint8 converts FlagRegister values to uint8
// every bit corresponds to a specific state, as described in FlagRegister doc (in a byte, the first bit is 7 and the
// last is 0)
func (f *FlagRegister) ToUint8() uint8 {
	// 1000 0000
	if f.zero {
		return 1 << ZeroFlagBytePosition
	}
	// 0100 0000
	if f.subtract {
		return 1 << SubtractFlagBytePosition
	}
	// 0010 0000
	if f.halfCarry {
		return 1 << HalfCarryFlagBytePosition
	}
	// 0001 0000
	if f.carry {
		return 1 << CarryFlagBytePosition
	}
	return uint8(0)
}

// ParseUint8 parses a uint8 into a FlagRegister
func (f *FlagRegister) ParseUint8(val uint8) {
	// 1000 0000
	if (val >> ZeroFlagBytePosition & 0b1) != 0 {
		f.zero = true
	}
	// 0100 0000
	if (val >> SubtractFlagBytePosition & 0b1) != 0 {
		f.subtract = true
	}
	// 0010 0000
	if (val >> HalfCarryFlagBytePosition & 0b1) != 0 {
		f.halfCarry = true
	}
	// 0001 0000
	if (val >> CarryFlagBytePosition & 0b1) != 0 {
		f.carry = true
	}
}
