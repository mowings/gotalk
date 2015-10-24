package main

import (
	"fmt"
)

const NUM_ROUTINES = 2000000

func main() {
	fmt.Println("Starting...")
	start := make(chan int)
	finish := make(chan int)
	in := start
	for i := 1; i <= NUM_ROUTINES; i++ {
		out := make(chan int)
		if i == NUM_ROUTINES {
			out = finish
		}
		go func(in, out chan int) {
			n := <-in
			n += 1
			out <- n
		}(in, out)
		in = out
	}
	fmt.Printf("Waiting for goroutines to complete.\n")
	start <- 0
	n := <-finish
	fmt.Printf("n = %d\n", n)
}
