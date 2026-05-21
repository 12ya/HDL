package main

import "testing"

func TestParser(t *testing.T) {
	tests := map[string]string{
		"D=M+1":       "M+1",
		"D;JEQ":       "D",
		"AMD=D-1;JNE": "D-1",
		"0;JMP":       "0",
		"M=D":         "D",
		"D=A":         "A",
	}

	for input, expected := range tests {
		got := comp(input)
		if got != expected {
			t.Errorf("comp(%q) = %q, expected %q", input, got, expected)
		}
	}

}
