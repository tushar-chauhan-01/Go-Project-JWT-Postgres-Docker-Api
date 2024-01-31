package main

import "fmt"

func main() {
	fmt.Println("Welcone to control flow statements in golang")

	age := 18

	// if else statements
	if age < 18 {
		fmt.Println("Your are not allowed to vote.")
	} else {
		fmt.Println("Your are allowed to vote.")
	}

	if age%2 == 0 {
		fmt.Println("Its even")
	} else {
		fmt.Println("Its odd")
	}

	// switch case
	day := 5
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
		fallthrough // this means next case will be allowed too to be executed
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Enter valid number between 1 and 7.")
	}

}
