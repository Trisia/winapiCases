package main

import (
	"fmt"
	"net"
)

func main() {
	cidr, ipNet, _ := net.ParseCIDR("10.0.2.1/24")
	fake := net.CIDRMask(24, 32)
	fmt.Println(cidr, ipNet, fake)
}
