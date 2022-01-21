package main

import (
	"fmt"
)

//Share memory by communicating (using channels)
func main() {

	ch := make(chan int)
	//go writeData(ch)
	fmt.Println("Initiating the goroutine")
	go func() {
		fmt.Println("goroutine - starting")
		ch <- 100
		fmt.Println("goroutine - completing")
	}()
	fmt.Println("Initiating the channel read")
	result := <-ch
	fmt.Println(result)

}

func writeData(ch chan int) {
	ch <- 100
}
