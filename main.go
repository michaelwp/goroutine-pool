package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	goSemaphore(10, 3)
}

func goSemaphore(n int, s int) {
	var sem = make(chan int, s)
	var wg sync.WaitGroup

	defer close(sem)

	wg.Add(n)
	for i := 1; i <= n; i++ {
		go longRunningProcess(i, &wg, sem)
	}
	wg.Wait()
}

func randomNumbersOfTime() time.Duration {
	rand.Seed(time.Now().UnixNano())
	var timeDuration = time.Duration(rand.Intn(10))

	return timeDuration
}

func longRunningProcess(i int, wg *sync.WaitGroup, sem chan int) {
	defer wg.Done()

	sem <- 1

	start := time.Now()
	fmt.Printf("Process: %d start \n", i)

	// pause the process for random amount of time in second
	time.Sleep(randomNumbersOfTime() * time.Second)
	fmt.Printf("Process: %d finish in: %v \n\n", i, time.Since(start))

	<-sem
}
