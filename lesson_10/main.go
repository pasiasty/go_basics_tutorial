package main

import "fmt"

func printArray(a [5]int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("a[%v] = %v\n", i, a[i])
	}
}

func printSliceInfo(s []string) {
	fmt.Printf("s = %v len(s) = %v cap(s) = %v\n", s, len(s), cap(s))
}

func main() {
	var a [5]int

	a[0] = 1
	a[1] = 3

	printArray(a)

	b := [3]string{"a", "b", "c"}
	fmt.Println("b =", b)

	var s []string = []string{"a", "b", "c"}
	printSliceInfo(s)
	s = append(s, "b", "c")
	printSliceInfo(s)
	s = append(s, "b", "c")
	printSliceInfo(s)
	fmt.Println("s =", s)

	var empty []int

	if empty == nil {
		fmt.Println("empty is nil!")
	}

	bigSlice := make([]string, 1024)
	printSliceInfo(bigSlice)

	surface := [][]int{
		{0, 1, 0},
		{1, 0, 1},
		{0, 1, 0},
	}

	surface[1][1] = 7

	fmt.Println("surface", surface)

	iterable := []int{5, 13, 28}
	for i := 0; i < len(iterable); i++ {
		iterable[i] = i
		fmt.Printf("iterable[%v] = %v\n", i, iterable[i])
	}

	for i, el := range iterable {
		el = i + 10
		fmt.Printf("iterable[%v] = %v\n", i, el)
	}

	for i := range iterable {
		fmt.Printf("iterable[%v] = %v\n", i, iterable[i])
	}

	for _, el := range iterable {
		fmt.Println("el =", el)
	}

	fmt.Println("iterable", iterable)
}
