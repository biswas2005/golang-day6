package practice

import "fmt"

func FindNumber() {

	slice := []int{4, 6, 2, 3, 12, 32, 45, 1, 7}

	var num int
	fmt.Print("Enter a number to check: ")
	fmt.Scan(&num)
	found := false

	for i := 0; i < len(slice); i++ {
		if slice[i] == num {
			found = true
			break

		}

	}
	if found {
		fmt.Println(num, "Is present.")

	} else {
		fmt.Println(num, "Is not present.")
	}

}
