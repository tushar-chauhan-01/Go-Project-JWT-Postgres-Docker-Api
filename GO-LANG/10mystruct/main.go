package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	info := User{"John", 21, true}
	fmt.Println(info)
	fmt.Printf("%+v \n", info) // this prints in key value form // %v is used for array or map or struct
	// %+v for whole struct
	fmt.Printf("Name is %v and age is %v \n", info.Name, info.Age)
}

// all class have to start with capital letter as it will be public
type User struct {
	Name   string
	Age    int
	Status bool
}
