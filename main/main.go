package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v)
	}
}

func merge(a, b <-chan int) <-chan int {

	// this should take a and b and return a new channel which will send all values from both a and b
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(2)

	go output(a)
	go output(b)

	// Start a goroutine to close out once all the output goroutines are
	// done.  This must start after the wg.Add call.
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func asChan(vs ...int) <-chan int {

	// this should return a channel and send `vs` values randomly to it.
	var ch = make(chan int, len(vs))
	defer close(ch)

	// random number source generation
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)

	var iterator = vs
	// pick random values from given array
	for i := 0; i < len(vs); i++ {
		index := r.Intn(len(iterator))
		value := vs[index]

		iterator = append(iterator[:index], iterator[index+1:]...)

		// push value to channel
		ch <- value
	}

	return ch
}
