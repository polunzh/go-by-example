package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [5]int
	fmt.Println("emp:", a)

	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)
	fmt.Println("type:", reflect.TypeOf(b))

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
			fmt.Printf("2D[%d][%d]:%d\n", i, j, twoD[i][j])
		}
	}

	fmt.Println("2D:", twoD)
	fmt.Println("length of 2D:", len(twoD))
}
