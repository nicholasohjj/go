package main

import "fmt"

func delete(x *[]int, i int) {
	if i < 0 || i >= len(*x) {
		return
	}
	*x = append((*x)[:i], (*x)[i+1:]...)
}

func deleteSwap(x *[]int, i int) {
	if i < 0 || i >= len(*x) {
		return
	}
	(*x)[i], (*x)[len(*x)-1] = (*x)[len(*x)-1], (*x)[i]
	*x = (*x)[:len(*x)-1]
}

func main() {
	x := []int{1,2,3,4,5}
	delete(&x, 210)
	fmt.Println(x)
	deleteSwap(&x, 2)
	fmt.Println(x)
}


