package main

import "fmt"

func myCondition(x, y int) bool {
	return (x > 3) && (y < 7)
}

func myMultiply(x int) (int, bool) {
	if x < 10 {
		return x * 2, true
	}
	return 0, false
}

func main() {
	x := 2
	y := 6

	if myCondition(x, y) {
		fmt.Println("Hello!")
	} else {
		fmt.Println("No hello!")
	}

	if result, success := myMultiply(9); success {
		fmt.Println("Result: ", result)
	} else {
		fmt.Println("Failed!")
	}

	z := 10

	if z < 20 {
		if z > 10 {
			fmt.Println("Nice!")
		} else {
			fmt.Println("Great!")
		}
	} else {
		fmt.Println("Bad!!!")
	}

	if z < 20 {
		fmt.Println("Hey!")
	} else if z < 30 {
		fmt.Println("Ho!")
	} else {
		fmt.Println("Boo!")
	}
}
