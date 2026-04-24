package main

import "fmt"

func greet(name string, age int, city string) {
	fmt.Println("Hello,", name)
	fmt.Println("You are", age, "years old")
	fmt.Println("You are from ", city)
}

func main() {
	var name string
	var age int
	var city string
	fmt.Println("Enter your name?")
	fmt.Scanln(&name)

	fmt.Println("Enter your age?")
	fmt.Scanln(&age)

	fmt.Println("Enter your city?")
	fmt.Scanln(&city)

	greet(name, age, city)
}
