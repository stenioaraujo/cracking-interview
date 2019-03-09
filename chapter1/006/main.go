package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(compression("aabcccccaaa"))

	fmt.Println(compression("aaa"))
	fmt.Println(compression("aa"))
	fmt.Println(compression("aabb"))
	fmt.Println(compression(""))
	fmt.Println(compression("abcd"))
	fmt.Println(compression("abcdddddd"))
	fmt.Println(compression("aaAA"))
	fmt.Println(compression("aaaAAA"))
}

func compression(s string) string {
	newStringSlice := make([]string, len(s))

	lastRune := '-'
	var count int
	for _, r := range s {
		if lastRune == r || lastRune == '-' {
			count++
		} else {
			newStringSlice = append(newStringSlice, fmt.Sprintf("%s%d", string(lastRune), count))
			count = 1
		}
		lastRune = r
	}
	if count > 0 {
		newStringSlice = append(newStringSlice, fmt.Sprintf("%s%d", string(lastRune), count))
	}

	newString := strings.Join(newStringSlice, "")
	if len(newString) < len(s) {
		return newString
	}

	return s
}
