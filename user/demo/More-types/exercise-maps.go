package main
import (
	"golang.org/x/tour/wc"
	"strings"
)
func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	m := make(map[string]int)
	for i := range fields {
		_, ok := m[fields[i]]
		if ok {
			m[fields[i]]++
		} else {
			m[fields[i]] = 1
		}
	}
	return m
}
func main() {
	wc.Text(WordCount)
}