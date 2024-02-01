package main

import "fmt"

func main() {
	fmt.Println("Functions in Golang")
	greet()
	fmt.Println("sum =", add(4, 7))
	sum, message := addMultiple(2, 3, 4, 5, 6)
	fmt.Println(message, sum)
}

func greet() {
	fmt.Println("Namaste, Hola, Ciao, Hallo")
}

// we have to give signature (datatype) of function if we want to return some value
func add(num1 int, num2 int) int {
	fmt.Println("Adding -", num1, "and", num2)
	return (num1 + num2)
}

// whenw e dont know the limit to input
func addMultiple(values ...int) (int, string) {
	fmt.Println("Adding", values)
	sum := 0
	for _, val := range values {
		sum = sum + val
	}
	return sum, "SUM ="
}
