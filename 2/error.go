package main

import (
	"fmt"
)

func main() {
	willclose := make(chan complex64, 10)
	willclose <- -1
	willclose <- -1i
	<-willclose
	<-willclose
	read := <-willclose
	fmt.Println(read)
}
