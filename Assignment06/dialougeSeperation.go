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

	aliceChannel := make(chan string)
	bobChannel := make(chan string)

	seperator(conevrsation, &aliceChannel, &bobChannel)
}

func seperator(conevrsation string, aliceChannel *chan string, bobChannel *chan string) {
	startidx := 0

	for idx, char := range conevrsation {
		if char == 94 {
			break
		}
		if char != 36 && char != 35 {
			continue
		} else {
			if char == 36 {
				go alicePrinter(aliceChannel)
				*aliceChannel <- ("alice : " + string(conevrsation[startidx:idx]))
			} else {
				go bobPrinter(bobChannel)
				*bobChannel <- ("bob : " + string(conevrsation[startidx:idx]))
			}
			startidx = idx + 1
		}
	}
}

func alicePrinter(aliceChannel *chan string) {
	fmt.Println(<-*aliceChannel)
}

func bobPrinter(bobChannel *chan string) {
	fmt.Println(<-*bobChannel)
}
