package main

import (
	"fmt"
	"sync"
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

	stringRune := []rune(forwardString)

	length := len(stringRune)

	var wg sync.WaitGroup

	for idx := 0; idx < len(stringRune)/2; idx++ {
		wg.Add(1)
		go routineSwapper(&stringRune, idx, length-idx-1, &wg)
	}

	wg.Wait()
	fmt.Println(string(stringRune))
}

func routineSwapper(stringRune *[]rune, start int, end int, wg *sync.WaitGroup) {
	arr := *stringRune
	arr[start], arr[end] = arr[end], arr[start]
	wg.Done()
}
