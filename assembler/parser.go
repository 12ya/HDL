package main

import "strings"

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
