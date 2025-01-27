package main

import (
	"fmt"
)

func main() {
	fmt.Println("Index Based Slicing Program")
	fmt.Println("Enter the two numbers seperated spaces on next line")

	//Hardcoding the Values As per Instructions
	hardCoded := []string{"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}

	//Declaring and Scanning start and end indices
	var start, end int
	fmt.Scanf("%d", &start)
	fmt.Scanf("%d", &end)

	//Necessary Index Validation
	if start > len(hardCoded)-1 || end > len(hardCoded)-1 || start < 0 || end < 0 {
		fmt.Println("The Indices You've Entered are out of bound")
		return
	} else if start == end {
		fmt.Println("The Indices You've Entered are equal")
		return
	} else if end < start {
		fmt.Println("Starting index is less then ending index")
		return
	}

	//Printing the three slice segments
	fmt.Println(hardCoded[:start])
	fmt.Println(hardCoded[start:end])
	fmt.Println(hardCoded[end:])

}
