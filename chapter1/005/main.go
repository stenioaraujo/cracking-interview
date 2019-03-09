package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func main() {
	fmt.Println(oneEdit("pale", "ple"))
	fmt.Println(oneEdit("pales", "pale"))
	fmt.Println(oneEdit("pale", "bale"))
	fmt.Println(oneEdit("pale", "bake"))

	fmt.Println(oneEdit("abc", "ab"))
	fmt.Println(oneEdit("ab", "abc"))
	fmt.Println(oneEdit("", ""))
	fmt.Println(oneEdit("pale", "pae"))
	fmt.Println(oneEdit("pale", "pan"))
}

func oneEdit(s1, s2 string) bool {
	s1Length := utf8.RuneCountInString(s1)
	s2Length := utf8.RuneCountInString(s2)
	lengthDifference := math.Abs(float64(s1Length - s2Length))

	var i, j, diff int
	if lengthDifference == 0 {
		for i < len(s1) && j < len(s2) {
			r1, r1Size := utf8.DecodeRuneInString(s1[i:])
			r2, r2Size := utf8.DecodeRuneInString(s2[j:])

			if r1 != r2 {
				diff++
			}

			i += r1Size
			j += r2Size
		}

		return diff <= 1
	} else if lengthDifference == 1 {
		var longest, shortest string
		if s1Length > s2Length {
			longest, shortest = s1, s2
		} else {
			longest, shortest = s2, s1
		}

		for i < len(longest) && j < len(shortest) {
			r1, r1Size := utf8.DecodeRuneInString(longest[i:])
			r2, r2Size := utf8.DecodeRuneInString(shortest[j:])

			if r1 != r2 {
				diff++
				i += r1Size
			} else {
				i += r1Size
				j += r2Size
			}
		}

		return diff <= 1
	}

	return false
}
