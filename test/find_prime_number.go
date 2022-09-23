// A concurrent prime sieve

package main

import "fmt"

var counter int
var filterMap = make(map[chan int]int)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		fmt.Println("generate", i)
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in chan int, out chan int, prime int) {
	fmt.Println("channels: ", filterMap[in], filterMap[out])
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	counter = 0
	ch := make(chan int) // Create a new channel.
	filterMap[ch] = counter
	counter++
	go Generate(ch)      // Launch Generate goroutine.
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		filterMap[ch1] = counter
		counter++
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
