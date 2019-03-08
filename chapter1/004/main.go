package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(palindromePermutation("Tact Coa"))
	fmt.Println(palindromePermutation("arara"))
	fmt.Println(palindromePermutation("Abba"))
	fmt.Println(palindromePermutation("obo"))

	fmt.Println(palindromePermutation(""))
	fmt.Println(palindromePermutation("a"))

	fmt.Println(palindromePermutation("apex"))
	fmt.Println(palindromePermutation("stenio"))
	fmt.Println(palindromePermutation("concatenar"))
	fmt.Println(palindromePermutation("pararalelepipedo"))
}

func palindromePermutation(s string) bool {
	s = strings.Replace(s, " ", "", -1)
	s = strings.ToLower(s)

	gates := make(map[rune]bool)

	for _, r := range s {
		gates[r] = !gates[r]
	}

	open := 0
	for _, v := range gates {
		if v {
			open++
		}
	}

	if open > 1 {
		return false
	}

	return true
}
