package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sums := 0
	for _, num := range nums {
		sums += num
	}

	fmt.Println("the sum is", sums)

	for index, num := range nums {
		fmt.Printf("index: %d, num: %d\n", index, num)
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "abc" {
		fmt.Printf("index: %d, char: %c, ascii: %d\n", i, c, c)
	}
}
