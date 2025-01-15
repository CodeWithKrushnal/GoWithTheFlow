package main

import (
	"fmt"
)

func main() {
	fmt.Println("Integer to WeekDay Converter")
	fmt.Println("Kindly enter the Integer on next line")

	//Declaring and Initializing the Days of Week in a map
	days := map[int]string{1: "Monday", 2: "Tuesday", 3: "Wednesday", 4: "Thursday", 5: "Friday", 6: "Saturday", 7: "Sunday"}

	//Declaring and Scanning the Weekday From User
	var weekDay int
	_, e1 := fmt.Scanf("%d", &weekDay)

	//Validation and Result
	if e1 == nil {
		if weekDay > 0 && weekDay <= 7 {
			fmt.Println("The Day for the index", weekDay, "is", days[weekDay])
		} else {
			fmt.Println("Invalid Number Entered")
		}
	} else {
		fmt.Println("Invalid Input", e1)
	}
}
