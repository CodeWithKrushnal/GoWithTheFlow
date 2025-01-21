
# \< Go Scheduling, Channels, Goroutines \>

Friday 21/01/2025

### Session Guide: [Sagar Sonwane](mailto:sagar.sonwane@joshsoftware.com) 


## Made with ♥️ from [Krushnal Patil](mailto:krushnal.patil@joshsoftware.com)


# 1. Why not use time.Sleep to wait for goroutines?

## Usage of time.Sleep should be always avoided because of following errors:

1. time.Sleep is a blocking call which results in program to pause execution for a specific duration  
2. Such operation results in wastage of valuable CPU time in production application.  
3. Inefficient usage of system resources  
4. Unpredictable behavior in concurrent programs is observed.

## Alternative Synchronization Approaches to obtain similar functionalities are: 
1. Usage of Channels  
2. Usage of Mutex  
3. Usage of Waitgroups

# 2. How does golang schedular manage goroutines, which causes the sequence to be random?

The Go scheduler uses a technique called "round-robin scheduling" to manage goroutines. 

The sequence of goroutines being scheduled appears to be random because the Go scheduler uses a pseudorandom number generator to select the next goroutine to run. This is done to prevent a goroutine from monopolizing the CPU and to ensure that all goroutines get a fair share of the CPU time.

# 3. Buffered channel of capacity 1 is used to make sure only 1 go routine can access the critical section at a time. Write a code to justify this statement 

```golang
package main

import (
	"fmt"
)

func main() {
	buffChan := make(chan int, 1)
	criticalSection := 0
	fmt.Println(criticalSection)
	go firstFunc(&buffChan)
	go secondFunc(&buffChan)
	fmt.Println(criticalSection)
	criticalSectionWriter(&buffChan, &criticalSection)
	fmt.Println(criticalSection)
	criticalSectionWriter(&buffChan, &criticalSection)
	fmt.Println(criticalSection)
}

func firstFunc(buffChan *chan int) {
	*buffChan <- 1
}

func secondFunc(buffChan *chan int) {
	*buffChan <- 2
}

func criticalSectionWriter(buffChan *chan int, CriticalPtr *int) {
	*CriticalPtr = <-*buffChan
}
```

# 4. Find use cases for buffered and unbuffered channels 

## Usage of Buffered Channels:
1. Enforcing Sequence of the goroutine Calls  
2. Task Queue Management.  
3. Rate Limiting on parallel operations using shared Resources  
4. Control Access of Critical Section.  
5. Key advantage of buffered channels is it at least frees up the resources occupied by 1 goroutine while the other goroutine is in suspended state or yet to be called

## Usage of Unbuffered Channels 

1. Strict Synchronization  
2. Immediate Processing  
3. Avoiding Race Condition  
4. Direct Communication