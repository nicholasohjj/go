package main

import "fmt"

type IPAddr [4]byte

func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3])
}

func  main() {
	ip := IPAddr{1,2,3,4}
	fmt.Println(ip)
}
