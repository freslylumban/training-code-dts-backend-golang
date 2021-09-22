package main

import (
	"fmt"
	"sync"
	"time"
)

func showLastName(name string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("Worker is starting..")
	fmt.Println("My name is : ", name)
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Worker is done..")
}

func main() {
	// Spawn goroutine sejumlah 10 item
	// Handle spawning goroutine using WaitGroup

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go showLastName("Sipahutar", &wg)
	}

	wg.Wait()
}
