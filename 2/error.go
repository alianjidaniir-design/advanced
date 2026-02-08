package main

import (
	"fmt"
)

func printer(ch chan<- bool) {
	ch <- true
}

func write(c chan<- int, x int) {
	fmt.Println("1", x)
	c <- x
	fmt.Println("2", x)
}

func f2(out <-chan int, in chan<- int) {
	x := <-out
	fmt.Println("Read (saeed)", x)
	in <- x
}
func main() {

	c1 := make(chan bool)
	go printer(c1)

}
