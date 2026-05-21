package problemSolving

import (
	"fmt"
)

func main() {
	count := 0
	input := "hi how you"
	for i := 0; i < len(input); i++ {
		// fmt.Println(i)
		if string(input[i]) != "a" && string(input[i]) != "e" &&
			string(input[i]) != "i" && string(input[i]) != "o" && string(input[i]) != "u" && string(input[i]) != " " {
			count++
		}
	}

	fmt.Printf("notvowel=%d", count)
	//  fmt.Println(nottvowel)
}

// }
