package main

import (
	"fmt"
	"strings"
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

	destTable := map[string]string{
		"":    "000",
		"M":   "001",
		"D":   "010",
		"DM":  "011",
		"A":   "100",
		"AM":  "101",
		"AD":  "110",
		"AMD": "111",
	}

	jumpTable := map[string]string{
		"":    "000",
		"JGT": "001",
		"JEQ": "010",
		"JGE": "011",
		"JLT": "100",
		"JNE": "101",
		"JLE": "110",
		"JMP": "111",
	}

}

func cleanLine(line string) string {
	line, _, _ = strings.Cut(line, "//")
	line = strings.TrimSpace(line)
	line = strings.ReplaceAll(line, " ", "")
	return line
}

type Command int

const (
	L_COMMAND Command = iota
	A_COMMAND
	C_COMMAND
)

var commandMap = map[Command]string{
	L_COMMAND: "L_COMMAND",
	A_COMMAND: "A_COMMAND",
	C_COMMAND: "C_COMMAND",
}

func (c Command) String() string {
	return commandMap[c]
}

func commandType(cmd string) Command {
	if strings.HasPrefix(cmd, "(") {
		return L_COMMAND
	}
	if strings.HasPrefix(cmd, "@") {
		return A_COMMAND
	}
	return C_COMMAND
}

func symbol(cmd string) string {
	// if A_COMMAND: remove @
	// if L_COMMAND: remove ( and )

	switch commandType(cmd) {
	case L_COMMAND:
		cmd = strings.TrimPrefix(cmd, "(")
		return strings.TrimSuffix(cmd, ")")

	case A_COMMAND:
		return strings.TrimPrefix(cmd, "@")

	default:
		return ""
	}
}

func dest(cmd string) string {
	if destination, _, found := strings.Cut(cmd, "="); found {
		return destination
	}
	return ""
}

func jump(cmd string) string {
	if _, j, found := strings.Cut(cmd, ";"); found {
		return j
	}

	return ""
}

func comp(cmd string) string {
	before, _, _ := strings.Cut(cmd, ";")
	if _, after, found := strings.Cut(before, "="); found {
		return after
	}

	return before
}
