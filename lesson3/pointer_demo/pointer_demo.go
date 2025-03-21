package main

import (
	"fmt"
	"reflect"
)

func demoPointerTypes() {
	fmt.Println("Demonstrate pointer types:")

	var myInt int
	fmt.Println("pointer to int:", reflect.TypeOf(&myInt)) // Get a pointer to myInt and print the pointer's type

	var myFloat float64
	fmt.Println("pointer to float64:", reflect.TypeOf(&myFloat)) // get a pointer to myFloat and print the pointer's type

	var myBool bool
	fmt.Println("Pointer to Boolean:", reflect.TypeOf(&myBool)) // get a pointer to myBool and print the pointer's type
}

func demoPointerHolder() {
	fmt.Println("Demo pointer holder")

	var myInt int
	var myIntPointer *int // delare a variable that holds a pointer to an int
	myIntPointer = &myInt // assign a pointer to the variable
	fmt.Println("Value of myIntPointer:", myIntPointer)

	var myFloat float64
	var myFloatPointer *float64 // declare a variable that holds a pointer to a float64
	myFloatPointer = &myFloat   // assign a pointer to the variable
	fmt.Println("Value of myFloatPointer:", myFloatPointer)
}

func demoChangePointerValue() {
	fmt.Println("Demo updating value at pointer")

	myInt := 4
	fmt.Println("Initial value of myInt:", myInt)
	myIntPointer := &myInt
	*myIntPointer = 8
	fmt.Println("Updated value at myIntPointer:", *myIntPointer)
	fmt.Println("Updated value of myInt:", myInt)
}

func main() {
	demoPointerTypes()
	demoPointerHolder()
	demoChangePointerValue()
}
