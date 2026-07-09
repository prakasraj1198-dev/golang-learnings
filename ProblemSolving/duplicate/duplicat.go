package duplicate`

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 2, 2, 5, 3, 6, 8, 2, 3, 6}

	maps := map[int]int{}
	for i := 0; i < len(arr); i++ {
		count := 0
		for j := 0; j < len(arr); j++ {
			if arr[i] == arr[j] {
				count++
			}

		}
		maps[arr[i]] = count
		// 		   maps =arr[i]
		// fmt.Print(maps[i], " ")
	}

	// fmt.Printf()

	//    for k, v := range maps {

	for number := range maps {
		fmt.Printf("%v ", number)
	}

}



-=======================
package main

import "fmt"

func main() {
	arr := []int{1, 1, 2, 2, 5, 5, 11, 6, 6, 7, 8, 8, 6, 12, 2, 12, 20,33,33,20, 6, 16, 16, 9}
	maps := make(map[int]int)
	var result []int
	for i := 0; i < len(arr); i++ {
		if maps[arr[i]] == 0 {
			  maps[arr[i]]=1
			result = append(result, arr[i])
		}
	}
	fmt.Println(result)
}

// -==================================
// package duplicate

// import "fmt"

// func duplicat() {
// 	arr := []int{1, 2, 2, 3}
// 	temp := 0
// 	for i := 0; i < len(arr)-1; i++ {
// 		fmt.Println(arr)
// 	 	if arr[i]==arr[i+1]{

// 		}
// 	//    temp=arr[i]

// 		// fmt.Println(arr)
// 	}
// 	fmt.Println(arr)
// }
