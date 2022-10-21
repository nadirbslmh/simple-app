package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("simple calculator")

	var a, b int

	fmt.Print("enter first number: ")

	fmt.Scan(&a)

	fmt.Print("enter second number: ")

	fmt.Scan(&b)

	fmt.Println("Result: ", Add(a, b))
}
