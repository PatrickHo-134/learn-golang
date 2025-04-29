package main

import (
	"fmt"
)

// Common way of defining a struct
var myStruct struct {
	number float64
	word   string
	toggle bool
}

func section1() {
	fmt.Println("--Section 1---")
	myStruct.number = 3.14
	myStruct.word = "pie"
	myStruct.toggle = true
	fmt.Println("Accessing struct field myStruct.number:", myStruct.number)
	fmt.Println("Accessing struct field myStruct.word:", myStruct.word)
	fmt.Println("Accessing struct field myStruct.toggle:", myStruct.toggle)
}

// Type definitions allow you to create types of your own.
// They let you create a new defined type that’s based on an underlying type.
type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed float64
}

func section2() {
	fmt.Println("\n--Section 2--")
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name:", porsche.name)
	fmt.Println("Top speed:", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description:", bolts.description)
	fmt.Println("Count:", bolts.count)
}

// Using defined types with functions
func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func section3() {
	fmt.Println("\n--Section 3--")
	var windshield part
	windshield.description = "ABC windshield"
	windshield.count = 1
	showInfo(windshield)
}

// Modifying a struct using a function
type subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s *subscriber) { // take a pointer as a parameter
	fmt.Println("Name:", s.name)
	fmt.Println("Monthly rate:", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) *subscriber { // return a pointer
	var s subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return &s // return a pointer to a struct instead of the struct itself
}

func applyDiscount(s *subscriber) {
	s.rate = 4.99
}

func section4() {
	fmt.Println("\n--Section 4--")
	subscriber1 := defaultSubscriber("Pat Ho")
	// subscriber1 is no longer a struct, it's a struct pointer
	applyDiscount(subscriber1) //subscriber1 is already a pointer
	printInfo(subscriber1)

	fmt.Println("")
	subscriber2 := defaultSubscriber("Peter Parker")
	printInfo(subscriber2)
}

// Struct literals

func section5() {
	myCar := car{name: "Ferrari", topSpeed: 300}
	fmt.Println("\n--Section 5--")
	fmt.Println("Name:", myCar.name)
	fmt.Println("Top speed:", myCar.topSpeed)
}

// Adding a struct as a field on another type
type address struct {
	street     string
	city       string
	state      string
	postalCode string
}

type employee struct {
	name        string
	salary      string
	homeAddress address
}

func section6() {
	fmt.Println("\n--Section 6--")
	employee1 := employee{name: "Patrick", salary: "$5000"}
	employee1.homeAddress.street = "456 Elm St"
	employee1.homeAddress.city = "Sydney"
	employee1.homeAddress.state = "NSW"
	employee1.homeAddress.postalCode = "2207"

	fmt.Println("Employee name:", employee1.name)
	fmt.Println("Street:", employee1.homeAddress.street)
	fmt.Println("City:", employee1.homeAddress.city)
	fmt.Println("State:", employee1.homeAddress.city)
	fmt.Println("Postal Code:", employee1.homeAddress.postalCode)

}

// Embedding struct
type myEmployee struct {
	name string
	address
}

func section7() {
	fmt.Println("\n--Section 7--")
	firstEmployee := myEmployee{name: "Tony Stark"}
	firstEmployee.street = "15 Lynesta"
	firstEmployee.state = "NSW"
	firstEmployee.city = "Bexley"
	firstEmployee.postalCode = "2207"

	fmt.Println("Name:", firstEmployee.name)
	fmt.Println("Street:", firstEmployee.street)
	fmt.Println("City:", firstEmployee.city)
	fmt.Println("State:", firstEmployee.state)
}

func main() {
	// Accessing struct fields
	section1()

	// Defining types and structs
	section2()

	// Using defined types with functions
	section3()

	// Modifying a struct using a function
	section4()

	// Struct literals
	section5()

	// Setting a struct within another struct
	section6()

	// Embedded structs
	section7()
}
