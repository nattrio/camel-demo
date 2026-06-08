package main

import "fmt"

func main() {
	input := "Hello world 123-1"
	output := ToCamelCase(input)
	fmt.Printf("Input:  %q\n", input)
	fmt.Printf("Output: %q\n", output)
}
