package main

import "fmt"

var name string
var age int

func main() {
	fmt.Println("Hello, world!")
	name = getName()
	age = getAge()
	fmt.Println(name)
	fmt.Println(age)
	loop01()
}

func getName() string{
	return "Tom"
}

func getAge() int{
	return 25
}

func loop01() {
	for i:=0; i<5; i++{
		fmt.Println(i)
	}
}