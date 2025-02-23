package find_odd_int

import "fmt"

func FindOdd(seq []int) int {
	result := 0

	for _, value := range seq {
		fmt.Println("value: ", value)
		fmt.Println("result before: ", result)
		result ^= value
		fmt.Println("result after: ", result)
	}

	return result
}
