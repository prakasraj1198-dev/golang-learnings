package Problemsolving

var result int

func searchInserts(num []int, target int) {
	for i := 0; i < len(num); i++ {
		if num[i] >= target {
			result = i

			break
		}

		// return i
		// }
		// return len(num)
	}
	// fmt.Println(num)

}

// func main() {
// 	num := []int{1, 3, 5, 6}
// 	// fmt.Println(searchInsert(num, 2))
// 	searchInsert(num, 2)

// 	fmt.Println(result)
// }

// -----=---=
// package main

// import "fmt"

// var result int

// func searchInsert(num []int, target int)  {
// 	for i := 0; i < len(num); i++ {
// 		if num[i] >= target {
// 			return i

// 		}

// 		// return len(num)
// 	}
// 	// fmt.Println(num)
// 	return len (num)
// }

// func main() {
// 	num := []int{1, 3, 5, 6}
// 	// fmt.Println(searchInsert(num, 2))
// 	searchInsert(num, 2)

// 	fmt.Println(result)
// }
