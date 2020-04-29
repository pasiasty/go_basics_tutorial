package main

import "fmt"

func stillOk(x, y int) bool {
	return x*y < 6
}

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i, "hello!")
	}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(i, j)
		}
	}

	x, y := 0, 0
	for stillOk(x, y) {
		fmt.Println(x, y, "hey")
		x++
		if x == 3 {
			y++
			x = 0
		}
	}

	counter := 0
	for {
		fmt.Println(counter, "hey")
		counter++
		if counter == 5 {
			break
		}
	}

	for i := 0; i < 6; i++ {
		fmt.Println(i, "Hey!")
		if i%2 == 0 {
			continue
		}
		fmt.Println(i, "You!")
	}
}
