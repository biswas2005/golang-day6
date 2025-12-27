package arrays

import "fmt"

func SumOfArrays() {
	total := 0
	arr := [3]int{1, 2, 3}
	for i := 0; i <= 2; i++ {

		total += arr[i]
	}
	fmt.Println(total)

}
