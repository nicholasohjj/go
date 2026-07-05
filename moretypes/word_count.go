package main

import (
"fmt"
"strings"
)

func WordCount(s string) map[string]int {
	res := make(map[string]int)
	words := strings.Fields(s)

	for _, v := range words {
		res[v] += 1
	}

	return res
}

func main() {
	s := "hello world hello"
	res := WordCount(s)
	fmt.Println(res)
}
