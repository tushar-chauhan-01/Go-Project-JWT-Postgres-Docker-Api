package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang.")

	var fruitList [4]string // index start with 0
	fruitList[0] = "apple"
	fruitList[2] = "mango"
	fruitList[3] = "ananas"
	fmt.Println("fruit list is :", fruitList)             // for no value at 1 index , go will print space
	fmt.Println("len of fruit list is :", len(fruitList)) // len will always give full length of list even though 1 index is empty
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	var vegList = [3]string{"potato", "carrot", "corriander"}
	fmt.Println("veg list is :", vegList)
	fmt.Println("len of veg list is :", len(vegList))
	fmt.Printf("Type of veg list is %T\n", vegList)
}
