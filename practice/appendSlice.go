package practice

import "fmt"

func Merge() {
	//Normal append
	slice0 := []int{1, 2, 3}
	slice0 = append(slice0, 4, 5, 6)
	fmt.Println(slice0)

	//If the new set of values are to be added infront of the slice
	slice := []int{1, 2, 3}

	slice = append([]int{4, 5, 6, 7, 8}, slice...)
	fmt.Println(slice)

	//If there exists 3 or more slices

	Ogslice := []int{4, 5, 6}
	newFront := []int{1, 2, 3}
	newBack := []int{7, 8, 9}

	Ogslice = append(newFront, append(Ogslice, newBack...)...)
	fmt.Println(Ogslice)
}
