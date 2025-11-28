package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
	c2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "saeed"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout")
	}

	c3 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		c3 <- "sea"
	}()

	select {
	case res := <-c3:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}
