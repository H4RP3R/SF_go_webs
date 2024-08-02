package main

import (
	"fmt"
	"testing"
)

func Test_rleEncode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"abbccc", "1a2b3c"},
		{"x", "1x"},
		{"", ""},
		{"xxtttf", "2x3t1f"},
		{"1122233", "213223"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s => %s", tt.input, tt.want)
		t.Run(name, func(t *testing.T) {
			if got := rleEncode(tt.input); got != tt.want {
				t.Errorf("rleEncode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rleDecode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"1a2b3c", "abbccc"},
		{"1x", "x"},
		{"", ""},
		{"2x3t1f", "xxtttf"},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%s => %s", tt.input, tt.want)
		t.Run(name, func(t *testing.T) {
			if got := rleDecode(tt.input); got != tt.want {
				t.Errorf("rleDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
