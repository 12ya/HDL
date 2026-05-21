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

	// destTable := map[string]string{
	// 	"":    "000",
	// 	"M":   "001",
	// 	"D":   "010",
	// 	"DM":  "011",
	// 	"A":   "100",
	// 	"AM":  "101",
	// 	"AD":  "110",
	// 	"AMD": "111",
	// }

	// jumpTable := map[string]string{
	// 	"":    "000",
	// 	"JGT": "001",
	// 	"JEQ": "010",
	// 	"JGE": "011",
	// 	"JLT": "100",
	// 	"JNE": "101",
	// 	"JLE": "110",
	// 	"JMP": "111",
	// }

}
