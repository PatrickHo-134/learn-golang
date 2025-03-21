# Lesson 3
This section covers Chapter 3 of Head First Go

## Content covered
- Declaring functions
- Functions and variable scope
- Error values
- Function parameters receive copies of the arguments
- Pointers
- Pointer types
- Using pointers with functions



## Key notes:
- The fmt.Printf and fmt.Sprintf functions format values they’re given. The first argument should be a formatting string containing verbs (`%d`, `%f`, `%s`, etc.) that values will be substituted for.

- Within a formatting verb, you can include a width: a minimum number of characters the formatted value will take up. For example, `%12s` results in a 12-character string (padded with spaces), %2d results in a 2-character integer, and %.3f results in a floating-point number rounded to 3 decimal places.

- If you want calls to your function to accept arguments, you must declare one or more parameters, including types for each, in the function declaration. The number and type of arguments must always match the number and type of parameters, or you’ll get a compile error.

- If you want your function to return one or more values, you must declare the return value types in the function declaration.

- You can’t access a variable declared within a function outside that function. But you can access a variable declared outside a function (usually at the package level) within that function.

- When a function returns multiple values, the last value usually has a type of error. Error values have an Error() method that returns a string describing the error.

- By convention, functions return an error value of nil to indicate there are no errors.

- You can access the value a pointer holds by putting a * right before it: *myPointer

- If a function receives a pointer as a parameter, and it updates the value at that pointer, then the updated value will still be visible outside the function.