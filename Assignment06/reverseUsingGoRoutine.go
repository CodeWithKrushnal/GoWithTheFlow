package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("String Reversal Using Goroutine")
	fmt.Println("Kindly Enter the String for reversal on next line")

	var forwardString string
	_, isEnteredError := fmt.Scanf("%s", &forwardString)

	if isEnteredError != nil {
		fmt.Println("You Haven't Enetered Anything, Exiting the Program")
		return
	}

	stringReversalchannel := make(chan string)
	go reverseString(forwardString, &stringReversalchannel)
	fmt.Println(<-stringReversalchannel)
	close(stringReversalchannel)
}

func reverseString(forwardString string, ch *chan string) {
	stringLength := len(forwardString)
	stringArray := make([]string, stringLength, stringLength)

	for idx, char := range forwardString {
		stringArray[stringLength-idx-1] = string(char)
	}
	*ch <- strings.Join(stringArray, "")
	return
}
