package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"
)

func f1(t int) {
	c1 := context.Background()
	c1, cancle := context.WithCancel(c1)
	defer cancle()
	go func() {
		time.Sleep(2 * time.Second)
		cancle()
	}()
	select {
	case <-c1.Done():
		fmt.Println("f1() done", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f1() timeout. r:", r)
	}
	return
}

func f2(t int) {
	c2 := context.Background()
	c2, cancle := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancle()

	go func() {
		time.Sleep(2 * time.Second)
		cancle()
	}()
	select {
	case <-c2.Done():
		fmt.Println("saeed() done", c2.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("saeed() timeout. r:", r)
	}

}

func f3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(2*t) * time.Second)
	c3, cancle := context.WithDeadline(c3, deadline)
	defer cancle()
	go func() {
		time.Sleep(2 * time.Second)
		cancle()
	}()
	select {
	case <-c3.Done():
		fmt.Println("f3() done", c3.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("f3() timeout. r:", r)
	}
	return
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
	fmt.Println("n", n)
	f1(n)
	f2(n)
	f3(n)
}
