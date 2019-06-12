package main
import "fmt"
func main() {
	m := make(map[string]int)
	m["Answer"] = 42 // 0
	fmt.Println("The value:", m["Answer"])
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // 0 也可能是不存在此key， 也可能是存在但是值为0
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}