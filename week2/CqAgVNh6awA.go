// this one is not finish yet
// reference: https://blog.golang.org/pipelines
package main

import (
	"fmt"
	"time"
)

func spread(main chan int, a chan int, b chan int, c chan int) {
	for {
		value := <-main
		a <- value
		b <- value
		c <- value
	}
}

func fanIn(out chan int) {
	var total int
	for num := range out {
		total += num
		fmt.Println("Total =", total)
	}
	fmt.Println("Finnal total =", total)
}

func printChannel(c chan int, out chan int, f func(int) int) {
	for {
		value := <-c
		x := f(value)
		out <- x
	}
}

func main() {
	main := make(chan int)
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	out := make(chan int)

	go spread(main, a, b, c)
	go printChannel(a, out, func(i int) int {
		return i * 2
	})
	go printChannel(b, out, func(i int) int {
		return i * 3
	})
	go printChannel(c, out, func(i int) int {
		return i * 4
	})
	go fanIn(out)
	for i := 1; i <= 1; i++ {
		main <- i
	}
	time.Sleep(3 * time.Second)
	close(out)
	time.Sleep(3 * time.Second)
}