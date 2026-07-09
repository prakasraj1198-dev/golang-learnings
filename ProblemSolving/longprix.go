package Problemsolving

import (
	"fmt"
)

func longestCommonPrefix(word []string) string {
	if len(word) == 0 {
		return " " //
	}
	base := word[0]
	// fmt.Println(base)
	// fmt.Println("word1", word[1])
	temp := "" // here store   value
	// fmt.Println("temp", temp)
	for i := 0; i < len(base); i++ {
		char := base[i] /// this line bycode type convertion
		// fmt.Println("char", char)
		// fmt.Println("base", base[i])

		for j := 0; j < len(word); j++ {
			// fmt.Println(char)
			if i >= len(word[j]) || word[j][i] != char {
				if i > len(word[j]) && word[i][j] != char {
					return temp // here  temp value return
					// fmt.Printf ("word [i]",int (word[j][i]))

				}
			}
			//  fmt.Println("char", char)

			temp = temp + string(char)

		}
	}
	return temp
}
func Prefix() {
	word1 := []string{"flower", "flow", "flight"}
	fmt.Println("Prefix:", longestCommonPrefix(word1))
	word2 := []string{"dog", "racecar", "car"}
	fmt.Println(longestCommonPrefix(word2))
}
