package main

import "fmt"

func main() {

	num := make(chan int, 5)
	counter := 10
	for i := 2; i < counter; i++ {
		select {
		case num <- i:
			fmt.Println(1)
		case num <- i * i:
			fmt.Println("proccessing", i)
		default:
			fmt.Println("Not")
		}
	}
	for {
		select {
		case num := <-num:
			fmt.Println("@", num, " ")
		default:
			fmt.Println("default")
			return
		}
	}

}
