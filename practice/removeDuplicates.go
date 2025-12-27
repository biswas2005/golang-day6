package practice

import "fmt"

func RemoveDuplicates() []int {
	slice := []int{1, 2, 3, 1, 4, 2, 5, 2, 4, 6}
	result := []int{}

	for _, n := range slice {
		exist := false
		for _, r := range result {
			if r == n {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, n)
		}

	}
	fmt.Println(result)
	return result

}
