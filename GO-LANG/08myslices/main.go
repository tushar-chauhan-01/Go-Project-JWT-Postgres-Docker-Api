package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var fruitList = []string{"apple", "mango", "ananas"} // syntax is same as list but we dont initialize the length for slices
	fmt.Printf("Type of fruit list is %T\n", fruitList)
	fmt.Println("fruit list is ", fruitList)
	fmt.Println("len of veg list is :", len(fruitList))

	fruitList = append(fruitList, "banana", "guava") // to add value in a slice
	fmt.Println("fruit list is ", fruitList)
	fmt.Println("len of veg list is :", len(fruitList))

	fmt.Println("fruit list is ", fruitList[1:3]) // slicing up list, all values from index 1 to 3
	// statting point is always inclusive and end point is non exculsive in range

	values := make([]int, 4)
	values[0] = 1
	values[1] = 3
	values[2] = 2
	values[3] = 999
	// values[4] = 5 // this will throw error as we initialized list to just 4 values
	fmt.Printf("Type of values list is %T\n", values)
	fmt.Println("values list is ", values)
	fmt.Println("len of values list is :", len(values))

	values = append(values, 111, 232, 222, 444)
	fmt.Println("values list is ", values)
	fmt.Println("Is values list sorted :", sort.IntsAreSorted(values))

	sort.Ints(values) //sorting int values
	fmt.Println("values list is ", values)

	// boolean value
	fmt.Println("Is values list sorted :", sort.IntsAreSorted(values))

	// how to remove value in slice based on index
	var courses = []string{"python", "go", "java", "node.js", "c"}
	fmt.Println("courses list is ", courses)

	index := 3

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("courses list is ", courses)
}
