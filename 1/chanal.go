package main

import "fmt"

//write to chanal and close it

func writeTochanal(c chan int, x int) {
	c <- x
	close(c)
}
func boolchanal(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}
func main() {
	var ch chan bool = make(chan bool)
	go boolchanal(ch, 7)
	for val := range ch {
		fmt.Println(val, "", 2)
	}
	fmt.Println("")
	for i := 0; i < 10; i++ {
		fmt.Println(<-ch, "")
	}
	fmt.Println()

	var c chan int = make(chan int)

	go writeTochanal(c, 5)
	for vay := range c {
		fmt.Println(vay)
	}
}
