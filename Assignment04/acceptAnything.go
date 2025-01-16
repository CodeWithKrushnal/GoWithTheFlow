package main

import (
	"fmt"
)

type Hello struct {
	message string
}

func main() {
	fmt.Println("AcceptAnything Program")
	fmt.Println(`Enter the Input Associated for the Follwoing Value Inputs
	1. -> Integer
	2. -> String
	3. -> Boolean
	4. -> Custom Datatype Hello`)

	var usersSlection int
	fmt.Scanln(&usersSlection)

	switch usersSlection {
	case 1:
		fmt.Println("You've Selected an option to Enter an Integer, Kindly Enter the Integer except 0 on next line")
		var userInput1 int
		fmt.Scanf("%d", &userInput1)
		for userInput1 == 0 {
			fmt.Println("Invalid Input please Enter the Integer value on next line")
			fmt.Scanf("%d", &userInput1)
		}
		acceptAnything(userInput1)

	case 2:
		fmt.Println("You've Selected an option to Enter a String, Kindly Enter the String on next line")
		var userInput2 string
		fmt.Scanln(&userInput2)
		for userInput2 == "" {
			fmt.Println("You Haven't Entered Anything, Kindly Enter the String on next line")
			fmt.Scanln(&userInput2)
		}
		acceptAnything(userInput2)

	case 3:
		fmt.Println("You've Selected an option to Enter a boolean value, Kindly Enter 1 for true and 2 for false on next line")
		var userInput3 int
		fmt.Scanf("%d", &userInput3)
		for userInput3 != 1 && userInput3 != 2 {
			fmt.Println("Invalid Input, Kindly Enter 1 for true and 2 for false on next line")
			fmt.Scanf("%d", &userInput3)
		}
		if userInput3 == 1 {
			acceptAnything(true)
		} else {
			acceptAnything(false)
		}

	case 4:
		fmt.Println("You've Selected an option to Enter a Custom Datatype value, Kindly Enter the Message on next line")
		var userInput4 Hello
		fmt.Scanln(&userInput4.message)
		for userInput4.message == "" {
			fmt.Println("You Haven't Entered Anything, Kindly Enter the message on next line")
			fmt.Scanln(&userInput4.message)
		}
		acceptAnything(userInput4)

	default:
		fmt.Println("Invalid Argument")
	}

}

func acceptAnything(userInput any) {
	switch userInput.(type) {
	case int:
		fmt.Println("The value entered is", userInput, "and its type is Integer")
	case string:
		fmt.Println("The value entered is", userInput, "and its type is String")
	case bool:
		fmt.Println("The value entered is", userInput, "and its type is Boolean")
	default:
		fmt.Println("The value entered is", userInput, "and its Custom type is Hello")
	}
}
