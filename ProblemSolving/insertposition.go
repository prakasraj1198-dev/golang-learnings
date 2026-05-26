package Problemsolving

import "fmt"

func searchInsert(num []int, target int) int {

	low := 0
	high := len(num) - 1
	mid := 0
	for low <= high {

		mid = (low + high) / 2
		if num[mid] == target {
			// fmt.Println("high valeu"high)
			return mid
		} else if num[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
func Position() {
	num := []int{1, 3, 5, 8, 10}

	// value1 := searchInsert(num,12)
	// fmt.Printf("num 11 insert position value=%d\n", value1)
	// fmt.Println(searchInsert(num, 2))
	fmt.Println("ans......", searchInsert(num, 11))
	fmt.Println(searchInsert(num, 11))
}

// -=-------------------------------  other way ------------------
// package insert

// import "fmt"

// func searchInsert(num []int, target int) int {

// 	low := 0
// 	high := len(num) - 1
// 	for low <= high {
// 		mid := (low + high) / 2
// 		if num[mid] == target {
// 			return mid
// 		} else if num[mid] > target {
// 			high = mid - 1
// 		} else {
// 			low = mid + 1
// 		}
// 	}

// 	return low
// }
// func Position() {
// 	num := []int{1, 3, 5, 6, 7}
// 	fmt.Println(searchInsert(num, 7))
// 	fmt.Println(searchInsert(num, 7))
// 	// value := searchInsert(num, 7)

// 	// fmt.Println(value)

// }
