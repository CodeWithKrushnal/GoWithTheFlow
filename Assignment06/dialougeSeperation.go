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
	exit := make(chan bool)

	seperator(conevrsation, aliceChannel, bobChannel, exit)
}

func seperator(conevrsation string, aliceChannel chan string, bobChannel chan string, exit chan bool) {
	startidx := 0

	go printer(aliceChannel, bobChannel, exit)
	for idx, char := range conevrsation {
		if string(char) == "^" { //Alternatively ASCII value 94 Could be Used for Checking
			exit <- true
			break
		}
		if string(char) != "$" && string(char) != "#" { //Alternatively ASCII values 36 & 35 respectively Could be Used for Checking
			continue
		}
		if string(char) == "$" {
			aliceChannel <- ("alice : " + string(conevrsation[startidx:idx]))
		} else {
			bobChannel <- ("bob : " + string(conevrsation[startidx:idx]))
		}
		startidx = idx + 1
	}
}

func printer(aliceChannel, bobChannel chan string, exit chan bool) {
	for true {
		select {
		case alice := <-aliceChannel:
			fmt.Println(alice)
		case bob := <-bobChannel:
			fmt.Println(bob)
		case <-exit:
			break
		}
	}
}
