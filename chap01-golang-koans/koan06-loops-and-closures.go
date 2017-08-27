package main

import (
	"fmt"
	"time"
)

func tally(value int, sum *int) {
	time.Sleep(10 * time.Millisecond)
	*sum = *sum + value // `*sum += value` works too
}

func main() {

	{
		// This is how tally() works:
		var sum int
		tally(1, &sum)
		tally(2, &sum)
		tally(3, &sum)
		tally(4, &sum)
		if sum != 10 {
			panic("Tally didn't work as expected!")
		}
	}

	{
		// This also works:
		var sum *int = new(int) // NOTE: the outer curly brackets create new scope, thus a new 'sum' variable.
		tally(1, sum)
		tally(2, sum)
		tally(3, sum)
		tally(4, sum)
		if *sum != 10 {
			panic("Tally didn't work as expected!")
		}
	}

	{
		// This is the same thing but wrapped in a Goroutine:
		myFunc := func() {
			var sum *int = new(int)
			tally(1, sum)
			tally(2, sum)
			tally(3, sum)
			tally(4, sum)
			if *sum != 10 {
				panic("Tally didn't work as expected!")
			}
		}
		go myFunc()
	}

	{
		// Which is identical to this:
		go func() {
			var sum *int = new(int)
			tally(1, sum)
			tally(2, sum)
			tally(3, sum)
			tally(4, sum)
			if *sum != 10 {
				panic("Tally didn't work as expected!")
			}
		}() // NOTE: notice the trailing `()`
	}

	{
		// Which behaves identically to this:
		var sum *int = new(int)
		go func() {
			tally(1, sum)
			tally(2, sum)
			tally(3, sum)
			tally(4, sum)
			if *sum != 10 {
				panic("Tally didn't work as expected!")
			}
		}()
	}

	{
		// Why does this fail?
		// How can it be fixed while keeping the same for-loop and goroutine structure?
		var sum *int = new(int)
		for i := 1; i <= 4; i++ {
			go func() {
				tally(i, sum)
			}()
		}
		time.Sleep(1 * time.Second) // Wait for goroutines to finish
		if *sum != 10 {
			panic(fmt.Sprintf("Tally didn't work as expected! Expected 10 but got %v", *sum))
		}
	}

}
