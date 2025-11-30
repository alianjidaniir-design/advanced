package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
	"time"
)

type Client struct {
	id      int
	integer int
}

type Result struct {
	job    Client
	square int
}

var size = runtime.GOMAXPROCS(0)
var client = make(chan Client, size)
var data = make(chan Result, size)

func worker(wg *sync.WaitGroup) {

	for c := range client {
		square := c.integer * c.integer
		output := Result{c, square}
		data <- output
		time.Sleep(time.Second)
	}
	wg.Done()

}
func create(n int) {
	for i := 0; i < n; i++ {
		c := Client{i, i}
		client <- c
	}
	close(client)
}
func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run workerpool.go ADDRESS:PORT")
		return
	}
	njobs, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	nworker, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	go create(njobs)

	finished := make(chan interface{})

	go func() {
		for d := range data {
			fmt.Printf("Client ID: %d\tint:", d.job.id)
			fmt.Printf("%d\tsquare: %d\n", d.job.integer, d.square)
		}
		finished <- true
	}()
	var wg sync.WaitGroup
	for i := 0; i < nworker; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(data)

	fmt.Printf("Finished: %v\n", <-finished)
}
