package practice

import "fmt"

func Practice() {

	//slice contains two types 'string' 'int'
	//map <string>--> <int>
	//string acts as an index(pointer) to int
	slice := map[string]int{
		"Abhi":    100,
		"Akash":   85,
		"Mujamil": 70,
		"Vishal":  20,
	}

	//To print value of a single index
	fmt.Println("Marks of Abhi:", slice["Abhi"])

	//delete keyword used to delete an existing value
	delete(slice, "Akash")
	fmt.Println("Marks of Akash:", slice["Akash"])

	//Here we don't have David in slice
	//Still will print with a value '0'
	fmt.Println("Marks of David:", slice["David"])

	//index stores the index
	//exist is a boolean type
	//exist shows if a record exists
	index, exist := slice["David"]
	fmt.Println("Marks of David :", index)
	fmt.Println("Does it exist:", exist)

}
