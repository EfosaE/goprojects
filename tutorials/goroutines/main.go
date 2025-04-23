package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	// "sync"
)

func workerOne(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 100 {
		fmt.Printf("Worker One: %d\n", i)
		time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)

	}
}

func workerTwo(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 100 {
		fmt.Printf("Worker Two: %d\n", i)
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

	}
}

// func main() {
// 	var wg sync.WaitGroup
// 	start := time.Now()

// 	wg.Add(1) // add one goroutine to wait for
// 	go workerOne(&wg)

// 	wg.Add(1) // add one goroutine to wait for
// 	go workerTwo(&wg)

// 	wg.Wait() // wait for all goroutines
// 	elapsed := time.Since(start).Seconds()
// 	fmt.Printf("All workers done in %.2f seconds\n", elapsed)
// }

func channelOne(channel chan int) {
	fmt.Println("Channel One sent info...")
	channel <- 1
	fmt.Println("Channel Info sent succesfully")
}

func channelTwo(channel chan int) {
	val := <-channel
	fmt.Println("Channel Two received:", val)
}

func main() {
	channel := make(chan int)

	go channelOne(channel)
	go channelTwo(channel)

	// Wait for goroutines to finish (for demo)
	time.Sleep(1 * time.Second)
}