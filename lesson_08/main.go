package main

import "fmt"

func modifyXY(x, y int) (int, int) {
	return x + y, x - y
}

func setXY(x, y *int) {
	*x = 3
	*y = 17
}

func initOrModify(ptr **int) {
	if *ptr == nil {
		*ptr = new(int)
		**ptr = 3
	} else {
		(*(*ptr))++
	}
}

func simpleSerialize(x int) string {
	return fmt.Sprintf("x value: %v", x)
}

func binarySerialize(x int) string {
	return fmt.Sprintf("x binary value: %08b", x)
}

func printSerialization(x int, serialize func(int) string) {
	fmt.Println("serialization:", serialize(x))
}

func main() {
	x := 12
	y := 36

	fmt.Println("x =", x, "y =", y)
	x, y = modifyXY(x, y)
	fmt.Println("x =", x, "y =", y)

	var xPtr *int
	fmt.Println("xPtr =", xPtr)
	fmt.Println("xPtr == nil=", xPtr == nil)
	xPtr = &x

	fmt.Println("xPtr =", xPtr)
	fmt.Println("*xPtr =", *xPtr)

	fmt.Println("x =", x)
	*xPtr = 13
	fmt.Println("x =", x)

	fmt.Println("x =", x, "y =", y)
	setXY(&x, &y)
	fmt.Println("x =", x, "y =", y)

	var ptr *int
	fmt.Println("ptr =", ptr)
	initOrModify(&ptr)
	fmt.Println("ptr =", ptr)
	fmt.Println("*ptr =", *ptr)
	initOrModify(&ptr)
	fmt.Println("*ptr =", *ptr)

	fmt.Println("simpleSerialize(13) = ", simpleSerialize(13))
	fmt.Println("binarySerialize(13) = ", binarySerialize(13))

	printSerialization(13, simpleSerialize)
	printSerialization(13, binarySerialize)
}
