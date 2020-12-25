package main

import (
	"fmt"
	"time"
)

var counter int

/*
Race condition happen when the outcome of two threads/go routines depends on the inter-leaving/switching between them.
In my case a function that increment a variable is executed twice in 2 go routines.
In this case the race condition can happen because one of them can read the variable and get blocked.
By the time it resumes the other one could have read and incremented the values thus making the blocked go routine
not increment the actual value.

*/
func main() {

	go increment(&counter, 1_000_000)
	go increment(&counter, 1_000_000)

	time.Sleep(time.Second)
	fmt.Printf("counter = %v\n", counter)
}

func increment(counter *int, times int) {
	for i := 0; i < times; i++ {
		*counter++
	}
}
