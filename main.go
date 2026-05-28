package main

func main() {

}

// ackage main

// import (
// 	"fmt"
// )
// func longestCommonPrefix(word []string) string {
// 	if len(word) == 0 {
// 		return " " //
// 	}
// 	base := word[0]
// 	// fmt.Println(base)
// 	// fmt.Println("word1", word[1])
// 	temp := "" // here store   value
// 	// fmt.Println("temp", temp)
// 	for i := 0; i < len(base); i++ {
// 		char := base[i] /// this line bycode type convertion
// 		// fmt.Println("char", char)
// 		// fmt.Println("base", base[i])

// 		for j := 0; j < len(word); j++ {
// 			// fmt.Println(char)
// 			if i >= len(word[j]) || word[j][i] != char {
// 				if i > len(word[j]) && word[i][j] != char {
// 					return temp // here  temp value return
// 					// fmt.Printf ("word [i]",int (word[j][i]))

// 				}
// 			}
// 			//  fmt.Println("char", char)

// 			temp = temp + string(char)

// 		}
// 	}
// 	return temp
// }
// func main() {
// 	word1 := []string{"flower", "flow", "flight"}
// 	fmt.Println("Prefix:", longestCommonPrefix(word1))
// 	// word2 := []string{"dog", "racecar", "car"}
// 	// fmt.Println(longestCommonPrefix(word2))
// 	fmt.Println.
// }
// // -=======================================-==next way
// package main

// import "fmt"

// func longestCommonPrefix(word []string) string {
// 	if len(word) == 0 {
// 		return " "
// 	}
// 	base := word[3]
// 	for i := 0; i < len(word); i++ {
// 		base := string(word[i])
// 		for j := 0; j < len(base); j++ {

// 			// fmt.Println(i,j)
// 			fmt.Println("BASE I", string(base[i]))
// 			fmt.Println("BASE J", string(base[j]))
// 			if base[i] == base[j] {
// 				// fmt.Println(string(base[i]),i, string(base[j]),j)
// 			}
// 			base2 := string(word[j])
// 			//  if base2(word[j]) == base(word[i]) {
// 			// if i >= len(word[j]) || word[j][i] == char {
// 			fmt.Printf("i value %d\n", i)
// 			fmt.Printf("j value %d\n", j)
// 			fmt.Println(base2)

// 			// return base[i:1]
// 			// return  base [i:2]

// 		}
// 		// fmt.Println("------------------")
// 	}

// 	return base[i:1]
// }

// func main() {
// 	words := []string{"flow", "flower", "flight"}
// 	// words := []string{"dog", "racecar", "car"}
// 	fmt.Println(longestCommonPrefix(words))
// }

// ---=---=-==--=
// package main

// import (
// 	"fmt"
// )

// func longestcommonprefix(word []string)  {
// if len(word) == 0 {
// return ""
// }
// base := word[0]
// temp:=0
// for i := 0; i < len(word); i++ {
// 	  if word[i]==word[i]
// 	  temp++
// char := base[i]
// for j := 0; j < len(word); j++ {
// if i >= len(word[j])
// base[i]<word[i] {
// }
// word[j][i] == char {

// return base[:1]
// }
// return  base[:2]
// }
// }
// return base
// }

// -=================================
// package main

// import (
// 	"fmt"
// )

// func longestCommonPrefix(strs []string) string {
// 	// if string == 0 {
// 	// 	return ""
// 	// }
// 	// left := 0
// 	// right := 0

// 	// for i := 0; i < len(word); i++ {
// 	// 	// if
// 	fmt.Println(i)

//	}
//
// // }
//
//	func main() {
//		word := []string{"flower", "flow", "flight"}
//		fmt.Println(longestCommonPrefix(word))
//	}
//
// -============================
// func main() {
// 	str := "apple,banana,cherry"
// 	// 1. Basic Split
// 	// parts := strings.Split(str, ",")
// 	for i := 0; i < len(str); i++ {
// 		fmt.Println(str)
// 		if str==<pars
// 		// fmt.Println(parts)
// 	}
// }

// -=========================================

// package main

// import (
// 	"fmt"
// )

// func longestCommonPrefix(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}

// 	base := strs[0]

// 	for i := 0; i < len(base); i++ {
// 		char := base[i]

// 		for j := 1; j < len(strs); j++ {
// 			if i >= len(strs[j]) || strs[j][i] != char {
// 				return base[:i]
// 			}
// 		}
// 	}

// 	return base
// }

// func main() {
// 	words := []string{"flower", "flow", "flight"}
// 	fmt.Println("Prefix:", longestCommonPrefix(words)) // Output: "fl"
// }
// ---=-=--==-
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func longestCommonPrefixSort(strs []string) string {
// 	if len(strs) == 0 {
// 		return ""
// 	}

// 	sort.Strings(strs)

// 	first := strs[0]
// 	last := strs[len(strs)-1]

// 	i := 0
// 	for i < len(first) && i < len(last) && first[i] == last[i] {
// 		i++
// 	}

// 	return first[:i]
// }
