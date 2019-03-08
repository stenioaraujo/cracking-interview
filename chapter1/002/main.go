package main

import "fmt"

func main() {
	fmt.Println(isPermutation("", ""))
	fmt.Println(isPermutation("abba", "aabb"))
	fmt.Println(isPermutation("stẽnio", "ẽinost"))
	fmt.Println(isPermutation("arara", "aarra"))

	fmt.Println(isPermutation("", "aarra"))
	fmt.Println(isPermutation("arara", "arera"))
}

func isPermutation(s1, s2 string) bool {
	runes := make(map[rune]int)

	for _, r := range s1 {
		runes[r]++
	}

	for _, r := range s2 {
		runes[r]--
	}

	for _, v := range runes {
		if v != 0 {
			return false
		}
	}

	return true
}
