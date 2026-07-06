// Approximate sqrt(x) with Newton's method: z -= (z*z - x) / (2*z). Iterate 10 times. Start with z := 1.0. Done when: sqrt(2) ≈ 1.4142. Then: loop until |z*z - x| < 1e-10 instead of 10 fixed iterations.

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %g", float64(e))
}

func sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for math.Abs(z*z - x) >= 1e-10 {
		z -= (z*z - x) / (2*z)
	}
	return z, nil
}


func main() {
	fmt.Println(sqrt(4))
	fmt.Println(sqrt(5))
	fmt.Println(sqrt(16))
	fmt.Println(sqrt(-1))
}
