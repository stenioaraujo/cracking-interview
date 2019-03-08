package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(hasUniqueWithMap("abc"))
	fmt.Println(hasUniqueWithMap("stẽnio"))
	fmt.Println(hasUniqueWithMap("abba"))

	fmt.Println(hasUniqueInPlace("abc"))
	fmt.Println(hasUniqueInPlace("stẽnio"))
	fmt.Println(hasUniqueInPlace("abba"))
}

func hasUniqueWithMap(s string) bool {
	runes := make(map[rune]bool)

	for _, r := range s {
		if runes[r] {
			return false
		}
		runes[r] = true
	}

	return true
}

func hasUniqueInPlace(s string) bool {
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		for j := i + size; j < len(s); {
			ra, sizeJ := utf8.DecodeRuneInString(s[j:])
			if r == ra {
				return false
			}
			j += sizeJ
		}
		i += size
	}
	return true
}
