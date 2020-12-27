package main

import (
	"fmt"
	"sync"
)

type ChopStick struct{ sync.Mutex }

type Philosopher struct {
	id              int
	leftCS, rightCS *ChopStick
}

func (philosopher *Philosopher) eat(wg *sync.WaitGroup) {
	for i := 0; ; i++ {
		philosopher.leftCS.Lock()
		philosopher.rightCS.Lock()

		fmt.Printf("Philosopher-%v - Eating round %v\n", philosopher.id, i)

		philosopher.leftCS.Unlock()
		philosopher.rightCS.Unlock()
	}

	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)

	philosophers := make([]*Philosopher, 5)
	chopSticks := make([]*ChopStick, 5)

	for i := 0; i < 5; i++ {
		chopSticks[i] = new(ChopStick)
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, chopSticks[i], chopSticks[(i+1)%5]}
	}

	for i := 0; i < 5; i++ {
		go philosophers[i].eat(&wg)
	}

	wg.Wait()
}
