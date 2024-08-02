package main

import (
	"fmt"
	"strings"
)

func checkAnagram(s1, s2 string) bool {
	s1 = strings.ReplaceAll(strings.ToLower(s1), " ", "")
	s2 = strings.ReplaceAll(strings.ToLower(s2), " ", "")

	if len(s1) != len(s2) {
		return false
	}

	for _, l1 := range s1 {
		for _, l2 := range s2 {
			if l1 == l2 {
				part1, part2, _ := strings.Cut(s2, string(l1))
				s2 = part1 + part2
				break
			}
		}
	}

	return len(s2) == 0
}

func main() {
	fmt.Println(checkAnagram("кабан", "банка"))
	fmt.Println(checkAnagram("кабан", "банан"))
}
