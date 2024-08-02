package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func rleEncode(input string) string {
	encoded := ""
	length := len(input)

	if length == 0 {
		return encoded
	}

	count := 1
	for i := 1; i < length; i++ {
		if input[i] == input[i-1] {
			count++
		} else {
			encoded += strconv.Itoa(count) + string(input[i-1])
			count = 1
		}
	}
	encoded += strconv.Itoa(count) + string(input[length-1])

	return encoded
}

func rleDecode(input string) string {
	decoded := ""
	length := len(input)

	if length == 0 {
		return decoded
	}

	countStr := ""
	for i := 0; i < length; i++ {
		if unicode.IsDigit(rune(input[i])) {
			countStr += string(input[i])
		} else {
			count, _ := strconv.Atoi(countStr)
			decoded += strings.Repeat(string(input[i]), count)
			countStr = ""
		}
	}

	return decoded
}

func main() {
	inputStr := "xxxyyiiii"
	countEncode := rleEncode(inputStr)
	fmt.Printf("input: %s, output: %s\n", inputStr, countEncode)

	inputStr = "1b7w9b7w"
	decodeCount := rleDecode(inputStr)
	fmt.Printf("input: %s, output: %s\n", inputStr, decodeCount)
}
