package main

import (
	"bufio"
	"fmt"
	"os"
)

func scan1() string {
	in := bufio.NewScanner(os.Stdin)
	in.Scan()

	if err := in.Err(); err != nil {
		fmt.Println("Input error:", err)
	}

	return in.Text()
}

func scan2() string {
	in := bufio.NewReader(os.Stdin)
	str, err := in.ReadString('\n')

	if err != nil {
		fmt.Println("Input error:", err)
	}

	return str
}

func main() {
	fmt.Println("Text:", scan2())
}
