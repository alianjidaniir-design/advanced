package main

import (
	"fmt"
	"sync"
	"time"
)

var Password *secret
var wg sync.WaitGroup

type secret struct {
	RMW      sync.RWMutex
	password string
}

func change2(pass string) {
	if Password == nil {
		fmt.Println("password is nil")
		return
	}
	fmt.Println("change2() function")
	Password.RMW.Lock()
	fmt.Println("change2() locked")
	time.Sleep(1 * time.Second)
	Password.password = pass
	Password.RMW.Unlock()
	fmt.Println("change2() unlocked")
}

func show() {
	defer wg.Done()
	defer Password.RMW.RUnlock()
	Password.RMW.RLock()
	fmt.Println("show function locked")
	time.Sleep(1 * time.Second)
	fmt.Println("pass value", Password.password)
}

func main() {

	Password = &secret{password: "1234"}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go show()
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		change2("234567")
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		change2("54321")
	}()
	wg.Wait()
	fmt.Println("Current password value:", Password.password)
}
