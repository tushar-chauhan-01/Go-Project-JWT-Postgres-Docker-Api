package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"monday", "tuesday", "wednesday", "thursday"}

	// continue - skips that iteration in loop
	// break - breaks complete loop

	// for loop normal
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	fmt.Println("---------------------------------")

	// for loop range
	for i := range days { // range traverse through all records on list
		fmt.Println(days[i])
	}
	fmt.Println("---------------------------------")

	// for each loop
	for index, day := range days {
		fmt.Println(index, " ", day)
	}
	fmt.Println("---------------------------------")

	num := 1
	for num < 10 {

		if num == 5 {
			num++
			continue
		}
		if num == 9 {
			goto label_method
		}
		fmt.Println("number is", num)
		num++
	}

label_method: // this is label method, used to go out of loop or break loop and execute if some condition is met
	fmt.Println("INSIDE GOTO LABEL")

}
