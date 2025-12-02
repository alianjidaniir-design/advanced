package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func Read() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Println(value)
		case readValue <- value:
		}
	}

}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("two")
		return
	}
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Going to create %d random numbers.\n", n)
	go monitor()
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			set(rand.Intn(n * n))
		}()
	}
	wg.Wait()
	fmt.Println("Last value", Read())
}
