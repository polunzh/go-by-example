package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42

	return &p
}

func main() {
	fmt.Println(person{"foo", 12})

	fmt.Println(&person{"foo", 12})

	fmt.Println(newPerson("zhang"))

	s := person{name: "zhenqiang", age: 24}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 25
	fmt.Println(s.age)
}
