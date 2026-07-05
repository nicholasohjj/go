// Write func swap(a, b *int) that swaps the values. Test in main. Done when: you understand why passing int wouldn't work.

package main

import "fmt"

func swap(a,b *int) {
	*a, *b = *b, *a
}
func main() {
	a, b := 1,2
	swap(&a, &b)
	fmt.Println(a,b)


}
