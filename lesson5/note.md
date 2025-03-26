# Lesson 5
This section covers Chapter 5, 6 and 7 of Head First Go

## 1. Arrays
- Zero values in arrays
- Array literals
- Accessing array elements within a loop
- Checking array length with the `len` function
- Looping over arrays safely with `for...range` loops
- Using the blank identifier with `for...range` loops

### Key notes:
* To declare an array variable, include the array length in square brackets and the type of elements it will hold: `var myArray [3]int`

* To assign or access an element of an array, provide its index in square brackets. Indexes start at 0, so the first element of `myArray` is `myArray[0]`.

* As with variables, the default value for all array elements is the zero value for that element’s type.

* You can set element values at the time an array is created using an array literal: `[3]int{4, 9, 6}`

* If you store an index that is not valid for an array in a variable, and then try to access an array element using that variable as an index, you will get a panic — a runtime error.

* You can get the number of elements in an array with the built-in `len` function.

* You can conveniently process all the elements of an array using the special `for...range` loop syntax, which loops through each element and assigns its index and value to variables you provide.

* When using a `for...range` loop, you can ignore either the index or value for each element by assigning it to the `_` blank identifier.

* The `os.Open` function opens a file. It returns a pointer to an `os.File` value representing that opened file.

* Passing an `os.File` value to `bufio.NewScanner` returns a `bufio.Scanner` value whose `Scan` and `Text` methods can be used to read a line at a time from the file as strings.

### Practice problem
- [Remove Duplicates from Sorted Array](https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/) (challenge 1)

### Additional reading material
- https://backend.anonystick.com/golang/basic-3.html

## 2. Slices
- Slice literals
- The slice operator
- Underlying arrays
- `append` function
- Slice and zero values
- Variadic functions
- Passing slices to variadic functions

### Key notes
* The type for a slice variable is declared just like the type for an array variable, except the length is omitted: `var mySlice []int`

* For the most part, code for working with slices is identical to code that works with arrays. This includes: accessing elements, using zero values, passing slices to the len function, and `for...range` loops.

* A slice literal looks just like an array literal, except the length is omitted: `[]int{1, 7, 10}`

* You can get a slice that contains elements `i` through `j - 1` of an array or slice using the slice operator: `s[i:j]`

* The `os.Args` package variable contains a slice of strings with the command-line arguments the current program was run with.

* A variadic function is one that can be called with a varying number of arguments.

* To declare a variadic function, place an ellipsis (`...`) before the type of the last parameter in the function declaration. That parameter will then receive all the variadic arguments as a slice.

* When calling a variadic function, you can use a slice in place of the variadic arguments by typing an ellipsis after the slice: `inRange(1, 10, mySlice...)`

### Practice problems
- [Merge Sorted Array](https://leetcode.com/problems/merge-sorted-array/description/)

### Additional reading material
- https://backend.anonystick.com/golang/basic-4.html

## 3. Maps
- Map literals
- Zero values within maps
- The zero value for a map variable is nil
- How to tell zero value apart from assigned values
- Removing key/value pairs with the `delete` function
- Using `for...range` loops with maps

### Key notes
* When declaring a map variable, you must provide the types for its keys and its values:
     ```
     var myMap map[string]int
     ```

* To create a new map, call the `make` function with the type of the map you want:
     ```
     myMap = make(map[string]int)
     ```

* To assign a value to a map, provide the key you want to assign it to in square brackets:
     ```
     myMap["my key"] = 12
     ```

* To get a value, you provide the key as well:
     ```
     fmt.Println(myMap["my key"])
     ```

* You can create a map and initialize it with data at the same time using a map literal:
     ```
     map[string]int{"a": 2, "b": 3}
     ```

* As with arrays and slices, if you access a map key that hasn’t been assigned a value, you’ll get a zero value back.

* Getting a value from a map can return a second, optional Boolean value that indicates whether that value was assigned, or if it represents a default zero value:
     ```
     value, ok := myMap["c"]
     ```

* If you only want to test whether a key has had a value assigned, you can ignore the actual value using the `_` blank identifier:
     ```
     _, ok := myMap["c"]
     ```

* You can delete keys and their corresponding values from a map using the delete built-in function:
     ```
     delete(myMap, "b")
     ```

* You can use `for...range` loops with maps, much like you can with arrays or slices. You provide one variable that will be assigned each key in turn, and a second variable that will be assigned each value in turn.
     ```
     for key, value := range myMap {
          fmt.Println(key, value)
     }
     ```

### Practice problems
- [Two Sum](https://leetcode.com/problems/two-sum/description/)

### Additional reading material
- https://backend.anonystick.com/golang/basic-6.html
