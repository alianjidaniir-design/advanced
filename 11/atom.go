package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type atomCounter struct {
	val int64
}

func (c *atomCounter) value() int64 {
	return atomic.LoadInt64(&c.val)
}

func main() {
	X := 123
	Y := 7
	var wg sync.WaitGroup
	counter := atomCounter{}
	for i := 0; i < X; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < Y; i++ {
				atomic.AddInt64(&counter.val, 1)
			}

		}()

	}
	wg.Wait()
	fmt.Println(counter.value())
}
