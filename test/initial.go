package main

import (
	"fmt"
	"time"
)

func say() {
	go func() {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("in sec go routine")
	}()

	fmt.Println("in first go routine")
}

func main() {
	go say()
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("in main go routine")
}
