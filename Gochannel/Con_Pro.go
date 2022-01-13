package Gochannel

import (
	"fmt"
)

const (
	BuffSize = 30
)

var size int = 0

func Produce(out chan<- int) {
	for number := 0; number < 10; number++ {
		out <- number
		size++
		fmt.Println("produce: ", number, "  size:  ", size)
	}
	close(out)
}

func Consumer(in <-chan int) {

	for data := range in {
		size--
		fmt.Println("consume: ", data, "  size:  ", size)
	}
}

func ProCon() {
	var ch = make(chan int, BuffSize)

	go Produce(ch)

	Consumer(ch)
}
