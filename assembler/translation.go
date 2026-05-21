package main

var destTable = map[string]string{
	"":    "000",
	"M":   "001",
	"D":   "010",
	"DM":  "011",
	"MD":  "011",
	"A":   "100",
	"AM":  "101",
	"MA":  "101",
	"AD":  "110",
	"DA":  "110",
	"AMD": "111",
	"ADM": "111",
	"MAD": "111",
	"MDA": "111",
	"DAM": "111",
	"DMA": "111",
}

var jumpTable = map[string]string{
	"":    "000",
	"JGT": "001",
	"JEQ": "010",
	"JGE": "011",
	"JLT": "100",
	"JNE": "101",
	"JLE": "110",
	"JMP": "111",
}

func destBits(mnemonic string) (string, bool) {
	bits, ok := destTable[mnemonic]
	return bits, ok
}

func jumpBits(mnemonic string) (string, bool) {
	bits, ok := jumpTable[mnemonic]
	return bits, ok
}
