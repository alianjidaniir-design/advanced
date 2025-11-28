package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var result = make(chan bool)

func timeout(t time.Duration) {

	temp := make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		defer close(temp)
	}()
	select {
	case <-temp:
		result <- false
	case <-time.After(t):
		result <- true
	}
}

func main() {
	ad := os.Args
	if len(ad) < 2 {
		fmt.Println("You need to provide a number!")
		return
	}
	f, err := strconv.Atoi(ad[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	Duration := time.Duration(f) * time.Millisecond
	fmt.Println("Duration:", Duration)
	go timeout(Duration)
	val := <-result
	if val {
		fmt.Println("Timeout!")

	} else {
		fmt.Println("OK")

	}

}
