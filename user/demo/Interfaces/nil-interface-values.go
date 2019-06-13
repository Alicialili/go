package main
import "fmt"
type I interface {
	M()
}
func main() {
	var i I
	describe(i)
	// i has no tuple to indicate which concrete method to call, will have runtime error
	i.M() // runtime error: invalid memory address or nil pointer dereference
}
func describe(i I) {
	fmt.Printf("(%+v, %T)\n", i, i)
}