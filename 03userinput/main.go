package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter you name - ")

	// comma ok/ error ok syntax -- input,err
	input, _ := reader.ReadString('\n') // \n in ReadString means that i will read the value till i get a new line
	fmt.Println("Thanks for input, ", input)
	fmt.Printf("Type of input is %T, ", input) // whatever you input it will always come in as string

}
