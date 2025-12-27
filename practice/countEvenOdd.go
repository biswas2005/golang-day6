package practice

import "fmt"

func Count() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	evenCount := 0
	oddCount := 0

	fmt.Println("Even numbers are: ")

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 {
			fmt.Println(slice[i])
			evenCount++

		}

	}
	fmt.Println("Odd numbers are: ")
	for i := 0; i < len(slice); i++ {

		if slice[i]%2 != 0 {
			fmt.Println(slice[i])
			oddCount++

		}
	}
	fmt.Println("Count of Even numbers", evenCount)
	fmt.Println("Count of Odd numbers", oddCount)

}
