package hightests

import "fmt"

func Hight() {
	arr := []int{10, 20, 6, 50, 60}
	var max int
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]

		}
	}
	fmt.Println(max)
}
