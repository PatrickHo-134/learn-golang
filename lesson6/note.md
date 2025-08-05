# Lesson 2
This section covers Chapter 8 and 9 of Head First Go

## Structs
- Slices and maps hold values of ONE type
- Structs are built out of values of MANY types
- Access struct fields using the dot operator
- Defined types and structs
- Using defined types with functions
- Modifying a struct using a function
- Accessing struct fields through a pointer
- Pass large structs using pointers
- A defined type's name must be capitalised to be exported
- Struct literals
- Adding a struct as a field on another type
- Setting up a struct within another struct
- Anonymous struct fields
- Embedding structs

### Key notes
- You can declare a variable with a struct type. To specify a struct type, use the `struct` keyword, followed by a list of field names and types within curly braces.

```
var myStruct struct {
   field1 string
   field2 int
}
```

- Writing struct types repeatedly can get tedious, so it’s usually best to define a type with an underlying struct type. Then the defined type can be used for variables, function parameters or return values, and so on.

```
type myType struct {
   field1 string
}
var myVar myType
```

- Struct fields are accessed via the dot operator.
```
myVar.field1 = "value"
fmt.Println(myVar.field1)
```

- If a function needs to modify a struct or if a struct is large, it should be passed to the function as a pointer.

- Types will only be exported from the package they’re defined in if their name begins with a capital letter.

- Likewise, struct fields will not be accessible outside their package unless their name is capitalized.

- Struct literals let you create a struct and set its fields at the same time.
```
myVar := myType{field1: "value"}
```

- Adding a struct field with no name, only a type, defines an anonymous field.

- An inner struct that is added as part of an outer struct using an anonymous field is said to be embedded within the outer struct.

- You can access the fields of an embedded struct as if they belong to the outer struct.

### Defined types
- Type errors in real life
- Defined types with underlying basic types
- Converting between types using functions
- Defining methods**
- The receiver parameters is (pretty much) just another parameter
- A method is (pretty much) just like a function
- Pointer receiver parameters


