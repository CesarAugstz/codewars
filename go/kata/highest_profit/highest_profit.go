package highest_profit

func MinMax(arr []int) [2]int {
	max := arr[0]
	min := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return [2]int{min, max}
}
