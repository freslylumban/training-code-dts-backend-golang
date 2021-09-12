package main

import (
	"fmt"
	"sync"
	"time"
)

func showName(name string, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("ShowName is starting...")
	fmt.Println("My name is ", name)

	time.Sleep(time.Second)

	fmt.Println("ShowName is done.")
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go showName("Fresly", &wg)
	}

	wg.Wait()
}
