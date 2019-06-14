package main
import (
	"fmt"
	"strings"
)
type IPAddr [4]byte
func (ip IPAddr) String() string {
	// return fmt.Sprintf("%v.%v.%v.%v", int(ip[0]), int(ip[1]), int(ip[2]), int(ip[3]))
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint([4]byte(ip))), "."), "[]")
}
func main() {
	hosts := map[string]IPAddr {
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}