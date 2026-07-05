// Print numbers 1–100. Multiples of 3 → Fizz. Multiples of 5 → Buzz. Multiples of both → FizzBuzz. Done when: works, uses switch (not chained if).

package main

import "fmt"

func main() {
	for i:=1; i < 101; i++ {
	switch {
		case i % 15 == 0:
			fmt.Println("FizzBuzz")
		case i % 5 == 0:
			fmt.Println("Buzz")
		case i % 3 == 0:
			fmt.Println("Fizz")
		default:
			fmt.Println(i)
	}
}
}
