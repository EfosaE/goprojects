package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func openConnection(done chan bool) {
	fmt.Println("Opening connection...")
	if num := rand.Intn(50); num < 25 {
		fmt.Println(num)
		time.Sleep(2 * time.Second)
	} else {
		fmt.Println(num)
		fmt.Println("Hanging Connection...")
		time.Sleep(2 * time.Hour)
	}

	done <- true

}

func openConnectionWithContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	done := make(chan bool)
	go openConnection(done)

	select {
		case <-done:
			fmt.Println("Connection established")
		case <-ctx.Done():
			fmt.Println("Connection timed out")
	}
}

func main() {
	openConnectionWithContext()
}
