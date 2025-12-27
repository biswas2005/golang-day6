package arrays

import "fmt"

func Maxmin() {
	arr := [5]int{4, 3, 2, 5, 1}

	max := arr[0]
	min := arr[0]

	for i := 0; i < len(arr); i++ {

		if arr[i] > max {
			max = arr[i]

		}

		if arr[i] < min {
			min = arr[i]
		}

	}
	fmt.Print(max, min)
}
