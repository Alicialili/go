package main
import "golang.org/x/tour/pic"
func Pic(dx, dy int) [][]uint8 {
	var ss [dy][dx]uint8
	for i, s := range ss {
		for ii := range s {
			s[ii] = uint8(i*ii)
		}
	}
	return ss
}
func main() {
	pic.Show(Pic(5, 5))
}