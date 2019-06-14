package main
import (
	"fmt"
)
func fibonacci(n int, c chan in) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c) // only the sender should close a channel, never the receiver
}
func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // receives values from channel repeatedly until it is closed
		fmt.Println(i)
	}
}