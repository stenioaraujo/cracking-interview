package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(isRotation("waterbottle", "erbottlewat"))

	fmt.Println(isRotation("stenio", "stenio"))
	fmt.Println(isRotation("stenioo", "ostenio"))
	fmt.Println(isRotation("stenioo", "enioest"))
}

func isRotation(original, rotation string) bool {
	if len(original) != len(rotation) {
		return false
	}

	return strings.Contains(rotation+rotation, original)
}
