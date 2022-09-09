package main 

import (
	"fmt"
	"strings"
)

var plaintext string = "This is quiz one. I hope I passed."

func encode(s string) string {
	s = strings.ToLower(s)
	var output string 
	for i := 0; i < len(s); i++ {
		n := int(s[i]) - 96
		if n > 0 && n < 27 {
			m := (2*n) % 26
			if m == 0 {
				m = 26
			}
			output = output + string(rune(m+96))
		}
	}
	return output
}

func main() {

	fmt.Println(encode("c"))
	fmt.Println(encode(plaintext))
}
