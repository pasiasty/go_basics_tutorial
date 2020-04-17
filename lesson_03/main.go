package main

import "fmt"

func main() {
	x := 1
	y := 2

	x = y

	y++ // y += 1
	x-- // x = x - 1 ; x -= 1

	x *= 3 // x = x * 3

	fmt.Printf("x = %v\n", x)
	fmt.Printf("y = %v\n", y)
	fmt.Printf("x + 1 - y = %v\n", (x+1)*y)
	fmt.Printf("5 %% 2 = %v\n", 5%2)
	fmt.Printf("5 / 2 = %v\n", 5./2)

	fmt.Printf("5 == 2 = %v\n", 5 == 2)
	fmt.Printf("2 >= 2 = %v\n", 2 >= 2)
	fmt.Printf("2 > 2 = %v\n", 2 > 2)
	fmt.Printf("2 != 2 = %v\n", 2 != 2)

	a := true
	b := false

	fmt.Printf("a && b = %v\n", a && b)
	fmt.Printf("a || b = %v\n", a || b)
	fmt.Printf("!a = %v\n", !a)

	n, k := 27, 30

	fmt.Printf("n = %v     (%06b)\n", n, n)
	fmt.Printf("k = %v     (%06b)\n", k, k)

	fmt.Printf("n | k = %v (%06b)\n", n|k, n|k)
	fmt.Printf("n & k = %v (%06b)\n", n&k, n&k)
	fmt.Printf("n ^ k = %v  (%06b)\n", n^k, n^k)

	fmt.Printf("n << 1 = %v (%06b)\n", n<<1, n<<1)
	fmt.Printf("n >> 3 = %v (%06b)\n", n>>3, n>>3)

	n |= k // n = n | k
}
