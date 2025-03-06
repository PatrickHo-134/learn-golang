package main

/////////////
// Section 1: first module

// import "fmt"

// func main() {
// 	fmt.Println("Hello, World!")
// }

/////////////
// Section 2: external packages

// import (
// 	"fmt"
// 	"rsc.io/quote"
// )

// func main() {
// 	fmt.Println(quote.Go())
// }

/////////////
// Section 3: variables and types

// import (
// 	"fmt"
// 	"math"
// 	"strings"
// )

// func main() {
// 	var quantity int
// 	var length, width float64
// 	var customerName string

// 	quantity = 4
// 	length, width = 64, 32
// 	customerName = "Patrick"

// 	fmt.Println(math.Floor(2.75))
// 	fmt.Println(strings.ToTitle("hello world"))

// 	fmt.Println(quantity)
// 	fmt.Println(length, width)
// 	fmt.Println(customerName)
// }

/////////////
// Section 4: Conversions

import (
	"fmt"
	"reflect"
)

func main() {
	var length float64 = 1.2
	var width int = 2

	length = float64(width)
	fmt.Println(length)
	fmt.Println(reflect.TypeOf(length))
}
