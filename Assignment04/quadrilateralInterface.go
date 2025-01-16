package main

import (
	"fmt"
)

type quadrilateral interface {
	calculateArea() float64
	calculatePerimeter() float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type square struct {
	side float64
}

func main() {
	fmt.Println("Program to Calculate the Area and Primeter of a Rectange and a Square with same interface")
	fmt.Println("Kindly Enter the length and Bredth of the rectangle on next line seperated by space on next line")

	var r rectangle
	fmt.Scanf("%f %f", &r.length, &r.breadth)
	for r.length <= 0 || r.breadth <= 0 {
		fmt.Println("You've Entered the incorrect length or breadth, Please enter it again")
		fmt.Scanf("%f %f", &r.length, &r.breadth)
	}

	fmt.Println("Kindly Enter the side of the square on next line")
	var s square
	fmt.Scanf("%f", &s.side)
	for s.side <= 0 {
		fmt.Println("You've Entered the incorrect value, Please enter it again")
		fmt.Scanf("%f", &s.side)
	}

	printAll(r)
	printAll(s)
}

func printAll(q quadrilateral) {
	fmt.Println("Quadrilateral Properties", q)
	fmt.Println("The Area of the Quadrilateral is:", q.calculateArea())
	fmt.Println("he Perimeter of the Quadrilateral is:", q.calculatePerimeter())
}

func (r rectangle) calculateArea() float64 {
	area := r.length * r.breadth
	return area
}

func (r rectangle) calculatePerimeter() float64 {
	perimeter := 2 * (r.length + r.breadth)
	return perimeter
}

func (s square) calculateArea() float64 {
	area := s.side * s.side
	return area
}

func (s square) calculatePerimeter() float64 {
	perimeter := 2 * (s.side + s.side)
	return perimeter
}
