package main

import (
	"fmt"
)

func main() {
	// 1. read .asm file
	// 2. clean lines
	// 3. pass 1: labels
	// 4. pass 2: translate
	// 5. write .hack file

	symbolTable := map[string]int{
		"SP":     0,
		"LCL":    1,
		"ARG":    2,
		"THIS":   3,
		"SCREEN": 16384,
		"KBD":    24576,
	}

	for i := 0; i <= 15; i++ {
		symbolTable[fmt.Sprintf("R%d", i)] = i
	}

}
