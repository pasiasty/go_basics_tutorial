package main

import "fmt"

func callIf(x int) {
	fmt.Println("Called for:", x)
	if x == 1 {
		fmt.Println("Hey!")
	} else if x == 2 {
		fmt.Println("Ho!")
	} else if x == 3 {
		fmt.Println("Boo!")
	} else if x == 4 {
		fmt.Println("Woohoo!")
	} else {
		fmt.Println("Invalid!")
	}
}

func callSwitch(x int) {
	fmt.Println("Called for:", x)
	switch x {
	case 1:
		fmt.Println("Hey!")
	case 2:
		fmt.Println("Ho!")
	case 3:
		fmt.Println("Boo!")
	case 4:
		fmt.Println("Woohoo!")
	default:
		fmt.Println("Invalid!")
	}
}

func getNumber() int {
	return 5
}

func main() {
	callIf(5)
	callSwitch(5)
	callIf(1)
	callSwitch(1)

	switch x := getNumber(); x * 2 {
	case 10:
		fmt.Println("Got 10")
	case 20:
		fmt.Println("Got 20")
	default:
		fmt.Println("Got something else", x)
	}

	y := 5
	switch x := getNumber(); x {
	case y:
		fmt.Println("Got y!")
	case 2 * y:
		fmt.Println("Got 2*y!")
	default:
		fmt.Println("Got something else", x, y)
	}

	z := 3
	switch {
	case z%2 == 0:
		fmt.Println("z is even")
	case z%2 == 1:
		fmt.Println("z is odd")
	}

	a := 2
	switch {
	case a < 5:
		fmt.Println("less then 5")
	case a < 10:
		fmt.Println("less then 10")
	default:
		fmt.Println("a too big")
	}
}
