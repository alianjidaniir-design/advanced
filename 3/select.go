package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

func gen(min, max int, createNumber chan int, end chan bool) {
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("Ended!")
		//return
		case <-time.After(time.Second * 4):
			fmt.Println("time.After()!")
			return
		}
	}
}

func main() {
	createNumber := make(chan int)
	end := make(chan bool)

	ars := os.Args
	if len(ars) != 2 {
		fmt.Println("please Use the a integer number")
		return
	}
	n, err := strconv.Atoi(ars[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		gen(0, n*n, createNumber, end)
		wg.Done()
	}()

	for i := 0; i < n+2; i++ {
		fmt.Printf("%d\n", <-createNumber)
	}
	end <- true
	wg.Wait()
	fmt.Println("\nExiting ...")
}
