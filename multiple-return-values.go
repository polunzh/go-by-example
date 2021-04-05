package main

import "fmt"

func vals() (int, int) {
	return 1, 2
}

func main() {
	v1, v2 := vals()
	fmt.Println("return values:", v1, v2)

	_, v := vals()
	fmt.Println("one of return value:", v)
}
