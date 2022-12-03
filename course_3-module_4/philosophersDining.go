/*
*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/
package main

import (
	"fmt"
	"sync"
)

var eatWgroup sync.WaitGroup

type chopstick struct{ sync.Mutex }

type philosopher struct {
	id                            int
	leftChopstick, rightChopstick *chopstick
}

func (p philosopher) eat() {
	defer eatWgroup.Done()
	p.leftChopstick.Lock()
	p.rightChopstick.Lock()
	printAction("started eating", p.id)
	p.rightChopstick.Unlock()
	p.leftChopstick.Unlock()
	printAction("finished eating", p.id)
}

func printAction(action string, id int) {
	fmt.Printf("Philosopher #%d %s\n", id+1, action)
}

func main() {
	// How many philosophers and chopsticks
	count := 5
	// Create chopsticks
	chopsticks := make([]*chopstick, count)
	for i := 0; i < count; i++ {
		chopsticks[i] = new(chopstick)
	}
	// Create philospohers slice
	philosophers := make([]*philosopher, count)
	// Assign chopsticks to philosophers
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			id: i, leftChopstick: chopsticks[i], rightChopstick: chopsticks[(i+1)%count]}
	}
	//Each philosopher calls eat() only 3 times
	for i := 0; i < 3; i++ {
		for _, v := range philosophers {
			eatWgroup.Add(1)
			go v.eat()
		}
	}
	eatWgroup.Wait()
}
