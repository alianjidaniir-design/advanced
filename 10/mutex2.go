package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

var m2 sync.Mutex
var v2 int

func changes() {
	m2.Lock()
	time.Sleep(1 * time.Second)
	v2 = v2 + 1
	if v2*v2 == 81 {
		v2 = 0
		fmt.Print("changed")

	}
}
func reads() int {
	m2.Lock()
	return 2
}

func main() {

	as := os.Args
	if len(as) < 2 {
		fmt.Println("We need to two arguments")
		return
	}
	mutex, err := strconv.Atoi(as[1])
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	var wg sync.WaitGroup
	fmt.Print(reads())
	for i := 0; i < mutex; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			changes()
			fmt.Printf("_> %d", reads())
		}()
	}

	wg.Wait()
	fmt.Printf("->%d\n", reads())
}
