## Lesson 1

This section covers Chapter 1 of Head First Go

### 1. Introduction and setting up
- Define your own package: `go mod init example/hello`
     + When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages. That go.mod file stays with your code, including in your source code repository.
- Call code in an external package: `go mod tidy`
     + When adding an external module as a requirement, a go.sum file is created mofor use in authenticating the module

### 2. Syntax basics
- Strings, runes, boolean, numbers
- Math operations
- Types
- Declaring variables
- Zero values: If you declare a variable without assigning it a value, that variable will contain the zero value for its type. For numeric types, the zero value is actually 0. But for other types, a value of 0 would be invalid, so the zero value for that type may be something else. The zero value for string variables is an empty string, for example, and the zero value for bool variables is false.
- Naming rules
     + If the name of a variable, function, or type begins with a capital letter, it is considered exported and can be accessed from packages outside the current one. (This is why the P in fmt.Println is capitalized: so it can be used from the main package or any other.) If a variable/function/type name begins with a lowercase letter, it is considered unexported and can only be accessed within the current package.
- Conversions

### 3. Compiling Go code
Using your favorite text editor, save our “Hello, Go!” code from earlier in a plain-text file named hello.go.

1. Open a new terminal or command prompt window.

2. In the terminal, change to the directory where you saved hello.go.

3. Run `go fmt hello.go` to clean up the code formatting. (This step isn’t required, but it’s a good idea anyway.)

4. Run `go build hello.go` to compile the source code. This will add an executable file to the current directory. On macOS or Linux, the executable will be named just hello. On Windows, the executable will be named hello.exe.

5. Run the executable file. On macOS or Linux, do this by typing `./hello` (which means “run a program named hello in the current directory”). On Windows, just type `hello.exe`.


## Sources:
- [Getting started](https://go.dev/doc/tutorial/getting-started)