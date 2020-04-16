package main

import "fmt"

var (
	s string
	b bool = true
	x int32
	y = 1.76
)

const c int = 16

func main() {

	x = int32(y)

	fmt.Printf("Hello x = %v y = %v b = %v s = '%v' c = %v!\n", x, y, b, s, c)
}
