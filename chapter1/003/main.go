package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(urlEncode("Mr John Smith    ", 13))
}

func urlEncode(s string, sentenceSize uint) string {
	bytes := make([]byte, len(s))
	perTwenty := []byte("%20")
	i := 0 // Walks the original string
	j := 0 // Walks the new string
	for k := uint(0); k < sentenceSize; k++ {
		r, size := utf8.DecodeRuneInString(s[i:])
		if r == ' ' {
			bytes = append(bytes[:j], perTwenty...)
			j += len(perTwenty)
		} else {
			c := string(r)
			bytes = append(bytes[:j], c...)
			j += size
		}
		i += size
	}
	return string(bytes)
}
