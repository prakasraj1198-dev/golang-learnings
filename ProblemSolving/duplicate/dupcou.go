package duplicate

import "fmt"

func Dupcou() {
	arr := []int{1, 1, 2, 2, 5, 3, 3, 6, 6, 7, 8, 8, 9}
	maps := make(map[int]int)

	// count := 0
	for i := 0; i < len(arr); i++ {

		maps[arr[i]]++
		// fmt.Printf("count: %v\n", count)
		// count++
	}

	//  fmt.Println(maps)

	for key, value := range maps {
		fmt.Println(key, value)
	}
	// 	for k := range maps {
	//  fmt.Printf("%v ", key)
}

// }
