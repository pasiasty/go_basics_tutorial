package main

import "fmt"

func printEmployeeInfo(name, lastName, position string) {
	fmt.Printf("Full name: %s %s\nPosition: %s\n\n", name, lastName, position)
}

func sumAndMultiply(x, y int) (int, int) {
	return x + y, x * y
}

func compute(x int) (result int) {
	result = 3
	return
}

func multipleReturn() int {
	return 1
	return 2
}

func main() {
	printEmployeeInfo("John", "Diggle", "Security")
	printEmployeeInfo("Oliver", "Queen", "CEO")
	printEmployeeInfo("Felicity", "Smoak", "CTO")

	x, y := 3, 5
	res1, res2 := sumAndMultiply(x, y)

	fmt.Printf("x = %v, y = %v, sum(x, y) = %v, %v\n", x, y, res1, res2)

	fmt.Printf("compute(x) = %v\n", compute(13))

	fmt.Printf("multipleReturn() = %v\n", multipleReturn())
	fmt.Printf("multipleReturn() = %v\n", multipleReturn())
}
