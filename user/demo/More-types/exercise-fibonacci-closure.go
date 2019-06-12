package main
import "fmt"
func fibonacci() func(int) int {
	return func(x int) int {
		if x < 2 {
			return x
		}
		return fibonacci()(x-2) + fibonacci()(x-1)
	}
}
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}