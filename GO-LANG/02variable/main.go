package main

import (
	"fmt"
)

var Token string = "jncijncowcw" // can be changed later in code

const Fixed_token = "jncijncowcw" // cant be changed ALSO ALL PUBLIC VARIABLE HAVE FIsDT LETTER AS CAPITAL

func main() {
	var name string = "John Doe"
	fmt.Println(name)                               // prints  a simple like not flexible in formatting a string
	fmt.Printf("%s is of type : %T \n", name, name) // prints with flexiblity of formatting a string
	fmt.Println("-------------------------------------")
	var islegal bool = true
	fmt.Println(islegal)
	fmt.Printf("islegal is of type : %T \n", islegal)
	fmt.Println("-------------------------------------")
	var age int = 25
	fmt.Println(age)
	fmt.Printf("%d is of type : %T \n", age, age)
	fmt.Println("-------------------------------------")
	var float_num float32 = 253.4353657896092897985
	fmt.Println(float_num)
	fmt.Printf("%.2f is of type : %T \n", float_num, float_num)
	fmt.Println("-------------------------------------")
	var float_num1 float64 = 253.4353657896092897985 // gives better precision than float 32
	fmt.Println(float_num1)
	fmt.Printf("%.2f is of type : %T \n", float_num1, float_num1)
	fmt.Println("-------------------------------------")

	// default values with alias
	var a int
	var b string
	var c bool
	fmt.Print(a, b, c)
	fmt.Println()
	fmt.Println("-------------------------------------")

	// variables without var AKA with only walrus operator , IT MUST BE USED ONLY INSIDE A METHOD,
	// TO USE OUTSIDE A METHOD OR WHILE CREATING A GLOBAL VARIABLE WE HAVE TO USE VAR VAR_NAME VAR_TYPE = VALUE
	animal := "Tiger"
	age_animal := 12
	fmt.Printf("%s is an animal whose average life span is of %d \n", animal, age_animal)
	fmt.Printf("animal is of type %T and age_animal is of type %T \n", animal, age_animal)
	fmt.Println("-------------------------------------")
	fmt.Println(Token)
	fmt.Println(Fixed_token)

}
