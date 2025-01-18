package main

import (
	"fmt"
)

func main() {
	fmt.Println("Program to Demonstrate the Panic and Recovery Mechanism")
	fmt.Println("This program is set to invoke and Recover panic Hence Enter valid indices to keep the program running and enter invalid indices to invoke a pannic")

	arr := []string{"One", "Two", "Three", "Four", "Five"}

	loopController := true

	//Enclosing the panic causing function and the recovery mechanism in a loop to make sure they keep running after the panic followed by recovery
	for loopController == true {
		fmt.Println("\n", arr)

		fmt.Println("\n Eneter a index for getting the array elements above\n")
		var idx int
		fmt.Scanf("%d", &idx)

		//Panic Causing Function
		func(arr []string, idx int) {

			//Panic Recovery and Loop controller mechanism
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered Error", r)
					fmt.Println("Press 1 to quit anything else to continue ")
					var temp int
					fmt.Scanf("%d", &temp)
					if temp == 1 {
						loopController = false
					}
					return
				}
			}()

			fmt.Println("The element at the entered index", idx, "is", arr[idx])
		}(arr, idx)
	}
}
