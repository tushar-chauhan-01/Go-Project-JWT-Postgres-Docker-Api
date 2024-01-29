/*
In Go, the package main declaration at the beginning of a file indicates that this Go file is intended to be an
executable program rather than a library. The main package is a special package in Go, and it serves as the entry point
for the executable. When you create a Go executable, it must have a main package, and the main function within that
package serves as the starting point of execution.
*/
package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers.")

	val := 2 // normal variable with initialization
	fmt.Printf("Normal variable is val = %d \n", val)
	fmt.Println("Normal variable is", "val = ", val)

	var ptr *int // creating pointer for a new variable
	fmt.Println("value of pointer is", ptr)

	var ptr1 = &val // & is used to reference the memory of valiable, her ptr1 is referencing address of val variables
	fmt.Println("value of pointer1 is", ptr1)
	fmt.Println("actual value of pointer1 is", *ptr1)

	*ptr1 = *ptr1 * 2
	fmt.Println("actual value of pointer1 is", *ptr1)
}
