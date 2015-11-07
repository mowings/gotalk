package main

import (
	"fmt"
	"os"
	"strconv"
)

const NUM_ROUTINES = 2000000

func main() {

	// Get worker count and passes from command line
	if len(os.Args) <= 2 {
		fmt.Println("Missing argument.")
		return
	}
	num_routines, _ := strconv.Atoi(os.Args[1])
	passes, _ := strconv.Atoi(os.Args[2])

	// Set up ring head and tail
	start := make(chan int)
	finish := make(chan int)
	in := start

	// Create goroutines and channels
	fmt.Printf("Creating %d goroutines\n", num_routines)
	for i := 1; i <= num_routines; i++ {
		out := make(chan int)
		if i == num_routines {
			out = finish
		}
		go func(in, out chan int) {
			for {
				n := <-in
				out <- n + 1
			}
		}(in, out)
		in = out
	}

	// Send token around ring
	fmt.Printf("Sending token around ring\n")
	n := 0
	for p := 0; p < passes; p++ {
		start <- n
		n = <-finish
	}

	fmt.Printf("n = %d\n", n)
}
