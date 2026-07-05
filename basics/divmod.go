// Write func divmod(a, b int) (int, int) returning quotient and remainder. Call it in main for (17, 5), (100, 7), (9, 3). Done when: uses multiple return values, handles b == 0 by returning (0, 0)

package main

import "fmt"

func divmod(a, b int) (int, int) {
	if b == 0 { return 0, 0 }

	return a/b, a%b 
}

func main() {
	fmt.Println(divmod(17,5))
	fmt.Println(divmod(100,7))
	fmt.Println(divmod(9,3))
	fmt.Println(divmod(9,0))
}
