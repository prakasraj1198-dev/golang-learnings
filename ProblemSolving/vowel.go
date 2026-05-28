package Problemsolving

import (
	"fmt"
)

func Vowel() {

	count := 0
	input := "madam"
	for i := 0; i < len(input); i++ {
		fmt.Println(i)
		if string(input[i]) == "a" || string(input[i]) == "e" || string(input[i]) == "i" ||
			string(input[i]) == "o" || string(input[i]) == "u" || string(input[i]) == "A" || string(input[i]) == "E" || string(input[i]) == "I" {
			count++

		}

	}
	fmt.Println(count)
}
