package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now().Hour()
	switch {
	case t < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmi := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("It's a bool")
		case int:
			fmt.Println("It's a int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmi(true)
	whatAmi(1)
	whatAmi("zhenqiang")
}
