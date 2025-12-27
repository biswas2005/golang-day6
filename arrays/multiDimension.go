package arrays

import "fmt"

func Multidimension() {
	array := [3][3]int{{3, 4, 5}, {1, 2, 3}, {4, 5, 9}}

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Print(array[i][j])

		}
		fmt.Println()
	}
}
