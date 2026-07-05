// Write a function that defers 5 print statements. Verify LIFO order. Done when: you can predict the output before running it.

package main

import "fmt"

func main() {
	for i:=0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
