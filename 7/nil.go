package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second * 2)
	for {
		select {
		case input := <-c:
			sum += input
		case <-t.C:
			c = nil
			fmt.Println(sum)
			wg.Done()
		}
	}
}

func send(c chan int) {
	fmt.Println()

	for {
		c <- rand.Intn(69)
	}
}

func main() {
	c := make(chan int)
	wg.Add(1)
	go add(c)
	go send(c)
	wg.Wait()
	fmt.Println()
	fmt.Println("done")
}
