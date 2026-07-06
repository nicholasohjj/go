package main

import (
"fmt"
"math"
)

type ErrNegative float64

func (e ErrNegative) Error() string {
	return fmt.Sprintf("Error: negative number %g", e)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNegative(f)
	}
	return math.Sqrt(f), nil
}

func main() {
	fmt.Println(Sqrt(-4))
	fmt.Println(Sqrt(4))
}
