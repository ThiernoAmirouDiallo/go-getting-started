package main

import (
	"fmt"
	"sync"
)

const EatingRoundsNumber = 3
const PhilosophersNumber = 5

type ChopStick struct{ sync.Mutex }

type EatRequest struct {
	philosopherId int
	replyChannel  *chan *chan int
}

type Philosopher struct {
	id              int
	leftCS, rightCS *ChopStick
	replyChannel    chan *chan int
}

func (philosopher *Philosopher) eat(c *chan EatRequest, wg *sync.WaitGroup) {
	for i := 0; i < EatingRoundsNumber; i++ {
		*c <- EatRequest{philosopher.id, &philosopher.replyChannel}

		doneChan := <-philosopher.replyChannel
		fmt.Printf("Philosopher-%v - Eating round %v\n", philosopher.id+1, i+1)

		/*philosopher.leftCS.Lock()
		philosopher.rightCS.Lock()

		philosopher.leftCS.Unlock()
		philosopher.rightCS.Unlock()
		*/

		fmt.Printf("Philosopher-%v - Done round %v\n", philosopher.id+1, i+1)
		*doneChan <- philosopher.id
	}

	wg.Done()
}

type RequestProcessor struct {
	eating   []bool
	waiting  map[int]EatRequest
	doneChan chan int
}

func (rp RequestProcessor) process(request EatRequest) {
	if rp.canEat(request.philosopherId) {
		rp.eating[request.philosopherId] = true
		*request.replyChannel <- &rp.doneChan
	} else {
		rp.waiting[request.philosopherId] = request
		rp.eating[request.philosopherId] = false
	}
}

func (rp RequestProcessor) canEat(id int) bool {
	return !rp.eating[leftIdx(id)] && !rp.eating[rightIdx(id)]
}

func (rp RequestProcessor) hasFinishedEating(idx int) {
	rp.eating[idx] = false

	for id, request := range rp.waiting {
		if rp.canEat(id) {
			delete(rp.waiting, request.philosopherId)
			rp.process(request)
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(PhilosophersNumber)

	c := make(chan EatRequest, PhilosophersNumber)
	philosophers := make([]*Philosopher, PhilosophersNumber)
	chopSticks := make([]*ChopStick, PhilosophersNumber)

	for i := 0; i < PhilosophersNumber; i++ {
		chopSticks[i] = new(ChopStick)
	}

	for i := 0; i < PhilosophersNumber; i++ {
		philosophers[i] = &Philosopher{i, chopSticks[i], chopSticks[(i+1)%PhilosophersNumber], make(chan *chan int)}
	}

	go runHost(&c)

	for i := 0; i < PhilosophersNumber; i++ {
		go philosophers[i].eat(&c, &wg)
	}

	wg.Wait()
}

func runHost(c *chan EatRequest) {
	rp := RequestProcessor{make([]bool, PhilosophersNumber), make(map[int]EatRequest), make(chan int)}

	for {
		select {
		case request := <-*c:
			rp.process(request)
		case id := <-rp.doneChan:
			rp.hasFinishedEating(id)
		}
	}
}

func leftIdx(idx int) int {
	return (idx - 1 + PhilosophersNumber) % PhilosophersNumber
}

func rightIdx(idx int) int {
	return (idx + 1) % PhilosophersNumber
}
