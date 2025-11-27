package main

import (
	"fmt"
	"sync"
)

//write to chanal and close it

func writeTochanal(c chan int, x int) {
	c <- x
	close(c)
}

func boolchanal(ch chan bool) {
	ch <- true
}
func main() {
	// create a buffer canal
	c := make(chan int)

	var waitGroup sync.WaitGroup
	waitGroup.Add(1)

	go func(c chan int) {
		defer waitGroup.Done()
		writeTochanal(c, 13)
		fmt.Println("Exit.")
	}(c)

	fmt.Println("Read:", <-c)

	_, ok := <-c
	if ok {
		fmt.Println("channel is open!")

	} else {
		fmt.Println("channel is close!")
	}

	waitGroup.Wait()

	var ch chan bool = make(chan bool)
	for i := 0; i < 6; i++ {
		go boolchanal(ch)
	}
	n := 0
	for v := range ch {
		fmt.Println(v)
		if v == true {
			n++
		}
		if n > 2 {
			fmt.Println("n:", n)
			close(ch)
			break
		}
	}
	for i := 0; i < 5; i++ {
		fmt.Println(<-ch)
	}

}
