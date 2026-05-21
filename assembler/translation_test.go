package main

import "testing"

func TestDestBits(t *testing.T) {
	tests := map[string]string{
		"":    "000",
		"M":   "001",
		"D":   "010",
		"DM":  "011",
		"A":   "100",
		"AM":  "101",
		"AD":  "110",
		"AMD": "111",
	}

	for input, expected := range tests {
		got, ok := destBits(input)
		if !ok {
			t.Fatalf("destBits(%q) not found", input)
		}
		if got != expected {
			t.Errorf("destBits(%q) = %q; want %q", input, got, expected)
		}
	}
}

func TestJumpBits(t *testing.T) {
	tests := map[string]string{
		"":    "000",
		"JGT": "001",
		"JEQ": "010",
		"JGE": "011",
		"JLT": "100",
		"JNE": "101",
		"JLE": "110",
		"JMP": "111",
	}

	for input, expected := range tests {
		got, ok := jumpBits(input)
		if !ok {
			t.Fatalf("jumpBits(%q) not found", input)
		}
		if got != expected {
			t.Errorf("jumpBits(%q) = %q; want %q", input, got, expected)
		}
	}
}
