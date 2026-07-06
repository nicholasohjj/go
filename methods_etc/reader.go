package main

import (
"fmt"
"io"
)

type Reader struct {
}

func (r Reader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

func main() {
	r := Reader{}

	b := make([]byte, 8)

	n, err := r.Read(b)
	
	if err == io.EOF {
		return
	}

	fmt.Printf("n = %v err = %v b = %q\n", n, err, b)
}
