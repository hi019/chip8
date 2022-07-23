package main

import "fmt"

func main() {
	reg := Registers{b: 0b1100100, c: 0b10010110}
	fmt.Printf("%016b", reg.GetBC())

}

// 1. convert	0000000001100100
// 2. shift		0110010000000000
//
// 3. or		0110010000000000
//			  - 0000000010010110
//            -------------------
//				0110010010010110
