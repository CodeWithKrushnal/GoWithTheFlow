# Program to Implement `sync.Mutex` to avoid parallal access of shared Resources

## Problem Statement

In the below code snippet concurrent goroutines execution corrupts a piece of data by accessing it simultaneously it leads in raise condition.

package main

import (
	"fmt"
	"time"
)

func isEven(n int) bool {
	return n%2 == 0
}

```go
func main() {
	n := 0

	go func() {
		nIsEven := isEven(n)
		time.Sleep(5 * time.Millisecond)
		if nIsEven {
			fmt.Println(n, " is even")
			return
		}
		fmt.Println(n, "is odd")
	}()

	go func() {
		n++
	}()

	// just waiting for the goroutines to finish before exiting
	time.Sleep(time.Second)
}
```

output when you run this : 1 is Even (may vary when you run multiple times)

Update the given code to print the correct output.
output once code is corrected: 0 is Even

