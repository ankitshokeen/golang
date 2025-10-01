package main

import (
	"fmt"
	"sync"
)

func syncWait() {
	fmt.Println("welcome to class of sync waitGroup")

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait()
}

func task(i int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("task %d started \n", i)
	// task processing
	fmt.Printf("task %d completed \n", i)
}
