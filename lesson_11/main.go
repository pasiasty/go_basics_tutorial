package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func readInterestingValues(a *int, f *float32, s *string) bool {
	success := true

	fmt.Println("Give me integer")
	if _, err := fmt.Scanln(a); err != nil {
		success = false
		fmt.Println("Failed to get a", err)
	}
	fmt.Println("Give me float")
	if _, err := fmt.Scanln(f); err != nil {
		success = false
		fmt.Println("Failed to get f", err)
	}
	fmt.Println("Give me string")
	if _, err := fmt.Scanln(s); err != nil {
		success = false
		fmt.Println("Failed to get s", err)
	}

	return success
}

func main() {
	var a int
	var f float32
	var s string

	success := readInterestingValues(&a, &f, &s)
	if !success {
		fmt.Println("Scanning went wrong!")
	} else {
		fmt.Println("Scanning succeeded!")
	}

	fmt.Println("a =", a)
	fmt.Println("f =", f)
	fmt.Println("s =", s)

	b, err := ioutil.ReadFile("input_file.txt")
	if err != nil {
		fmt.Println("Failed to read file", err)
		return
	}

	fmt.Println("read file:", string(b))

	js := 0

	for _, c := range string(b) {
		if c == 'j' {
			js++
		}
	}

	fmt.Println("Found js:", js)
	result := fmt.Sprintf("Found js: %v", js)

	if err := ioutil.WriteFile("output_file.txt", []byte(result), os.ModePerm); err != nil {
		fmt.Println("Failed to write file:", err)
		return
	}
}
