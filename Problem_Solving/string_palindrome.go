package problemSolving   

import (
	"fmt"
)

func isPalindrome(s string) {
	reversed := ""

	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
		fmt.Println(s[i])
		
	} 
		if s == reversed {
		fmt.Printf("%v this is palindrome", s)

	} else {
		fmt.Printf("%v this is  not Palindrome", s)
	}
}

func Palindrome() {
	isPalindrome("mom")

	isPalindrome("madam")
	isPalindrome("welcome")

}

// -===========================  two way---
// package main

// import "fmt"

// func isPalindrome(s string) bool {
// 	reversed := ""
// 	for i := len(s) - 1; i >= 0; i-- {
// 		reversed += string(s[i])
// 	}
// 	return s == reversed
// }

// func main(){
// 	word :="madam"
// 	if is palindrome (word){
// 		fmt.Println("yes this palindrome")
// 	}else{
// 		fmt.Println("this not palindrome")
// 	}
// }
