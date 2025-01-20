package main

import (
	"fmt"
)

func main() {
	fmt.Println("Dialouge Seperation from string")
	fmt.Println("Kindly Enter the conevrsation String on next line")

	var conevrsation string
	_, isEnteredError := fmt.Scanf("%s", &conevrsation)

	if isEnteredError != nil {
		fmt.Println("You Haven't Enetered Anything, Exiting the Program")
		return
	}

	startidx := 0

	for idx, char := range conevrsation {
		if char == 94 {
			break
		}
		if char != 36 && char != 35 {
			continue
		} else {
			if char == 36 {
				fmt.Println("alice : ", string(conevrsation[startidx:idx]))
			} else {
				fmt.Println("bob : ", string(conevrsation[startidx:idx]))
			}
			startidx = idx + 1
		}
	}
}
