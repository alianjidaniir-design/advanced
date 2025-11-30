package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(sig os.Signal) {
	fmt.Println("Received signal:", sig)
}
func main() {
	fmt.Printf("Process ID : %d\n", os.Getpid())
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	start := time.Now()
	go func() {
		for {
			sig := <-sigs
			switch sig {
			case syscall.SIGINT:
				duration := time.Since(start)
				fmt.Println("runtime :", duration)
			case syscall.SIGINFO:
				handleSignal(sig)
				os.Exit(0)
			default:
				fmt.Println("Received signal:", sig)
			}
		}
	}()

	for {
		fmt.Println("+")
		time.Sleep(3 * time.Second)

	}

}
