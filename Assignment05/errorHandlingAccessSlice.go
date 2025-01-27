package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Program to Demonstarte Error Craetion and Handling Mechanism")
	fmt.Println("This program is set to invoke and handle error Hence Enter valid indices to keep the program running and enter invalid indices to invoke an error")

	slc := []string{"One", "Two", "Three", "Four", "Five"}

	var idx int
	fmt.Scanf("%d", &idx)

	res, err := accessSlice(slc, idx)

	//Checking if Function returnes any error
	if err == nil {
		fmt.Println("The Element at index", idx, "is", res)
	} else {
		fmt.Println("Error Occured in function accessSlice", err)
	}
}

// Returnes Element present at given valid index, returnes error if invalid index is passed
func accessSlice(slc []string, idx int) (string, error) {
	if idx > len(slc)-1 || idx < 0 {
		return "", errors.New("Index Out of Bound")
	} else {
		return slc[idx], nil
	}
}
