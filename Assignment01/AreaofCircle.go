package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Circle Area Calculator")
	fmt.Println("Enter the radius of Circle on next Line")

	//Decalring the radius variable and Accepting the input from the user
	var radius float64
	fmt.Scan(&radius)

	area := math.Round((math.Pi * math.Pow(radius, 2)*100))/100

	fmt.Println("The area of the circle of radius", radius, "is", area)
}