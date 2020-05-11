package main

import "fmt"

func baz() {
	bar()
}

func bar() {
	foo()
}

func foo() {
	panic("terminated by user in foo")
}

func prepare1() {
	fmt.Println("prepare 1")
}

func deinit1() {
	fmt.Println("deinit 1")
}

func prepare2() {
	fmt.Println("prepare 2")
}

func deinit2() {
	fmt.Println("deinit 2")
}

func prepare3() {
	fmt.Println("prepare 3")
}

func deinit3() {
	fmt.Println("deinit 3")
}

func main() {
	// baz()

	// var aPtr *int
	// fmt.Println("*aPtr = %v", *aPtr)

	prepare1()
	defer deinit1()

	if true {
		prepare2()
		defer deinit2()
	}

	prepare3()
	defer deinit3()

	fmt.Println("doing important work")
}
