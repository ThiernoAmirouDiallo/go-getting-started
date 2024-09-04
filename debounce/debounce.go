package debounce

import (
	"fmt"
	"sync"
	"time"
)

type debouncer struct {
	waitTime time.Duration
	mutex    sync.Mutex
	timer    *time.Timer
}

var DebouncedPrint func(name string) = func(name string) {
	fmt.Println("Debounced function called with", name)
}

// Debounce is a function that takes a function and a duration and returns a function that will only call the original function once the duration has passed since the last call.
func Debounce(f func(name string), waitTime time.Duration) func(name string) {
	context := &debouncer{waitTime: waitTime}

	return func(name string) {
		context.add(name, f)
	}
}

func (c *debouncer) add(name string, f func(name string)) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.timer != nil {
		c.timer.Stop()
	}

	c.timer = time.AfterFunc(c.waitTime, func() {
		f(name)
	})
}

func Example() {
	debouncedFunc := Debounce(DebouncedPrint, 50*time.Microsecond)

	debouncedFunc("- first call")

	time.Sleep(100 * time.Microsecond)

	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			debouncedFunc(fmt.Sprintf("%d-%d", i, j))
		}
		time.Sleep(200 * time.Millisecond)
	}

	// debounce test
	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("t is set to 3s, function is called at 1s, 3s, and 4s")
	debouncedFunc1 := Debounce(DebouncedPrint, 3*time.Second)
	time.AfterFunc(1*time.Second, func() { debouncedFunc1("1(1)") })
	time.AfterFunc(3*time.Second, func() { debouncedFunc1("3(1)") })
	time.AfterFunc(4*time.Second, func() { debouncedFunc1("4(1)") })

	time.Sleep(10 * time.Second) // wait 2 seconds

	fmt.Println("---------------------------------------------------------------------------------------")
	fmt.Println("t is set to 3s, function is called at 1s, 3s, and 4s")
	debouncedFunc2 := Debounce(DebouncedPrint, 3*time.Second)
	time.AfterFunc(1*time.Second, func() { debouncedFunc2("1(2)") })
	time.AfterFunc(3*time.Second, func() { debouncedFunc2("3(2)") })
	time.AfterFunc(8*time.Second, func() { debouncedFunc2("4(2)") })

	time.Sleep(12 * time.Second) // wait 2 seconds
}
