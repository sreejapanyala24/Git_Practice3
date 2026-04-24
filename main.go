package main

import "fmt"

func greet(name string) {
	fmt.Println("Hello,", name)
}

func main() {
	var name string
	fmt.Println("Enter your name?")
	fmt.Scanln(&name)
	greet(name)
}
