package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program to Calculate the Area and Primeter of a Rectange")
	fmt.Println("Kindly Enter the length and Bredth of the rectangle on next line seperated by space on next line")

	var length, bredth float64
	fmt.Scanf("%f %f", &length, &bredth)
	for length <= 0 || bredth <= 0 {
		fmt.Println("You've Entered the incorrect length or breadth, Please enter it again")
		fmt.Scanf("%f %f", &length, &bredth)
	}

	fmt.Println("The Area of the Rectangle is:", calculateArea(length, bredth))
	fmt.Println("The Perimeter of the Rectangle is:", calculatePerimeter(length, bredth))
}

func calculateArea(length float64, breadth float64) float64 {
	area := length * breadth
	return area
}

func calculatePerimeter(length float64, breadth float64) float64 {
	perimeter := 2 * (length + breadth)
	return perimeter
}
