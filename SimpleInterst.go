package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Simple Interest Calculator")
	fmt.Println("Enter the Principle, Rate of Interst seperated by comma")

	//Declaring variable and Accepting the rawInput from user as a string containing the comma seperated numbers
	var rawInput string
	fmt.Scanln(&rawInput)

	//Splitting the comma seperated string rawInput into an array(slice) of the seperated numbers in string state
	values:=strings.Split(rawInput,",")

	//Declaring New Variables for the Operation
	var Principle, ROI, TII int

	Principle,e1:=strconv.Atoi(values[0])
	ROI,e2:=strconv.Atoi(values[1])
	TII,e3:=strconv.Atoi(values[2])

	if e1!=nil || e2 !=nil || e3!=nil{
		fmt.Println("Only Input Integer Value")
		return
	}

	SI := (Principle * ROI * TII) / 100
	fmt.Println("The Simple Interest generated on the given Principle $", Principle, "in", TII, "years is $", SI)
	fmt.Println("Therefore the Total amount to be paid is: $", Principle+SI)
}
