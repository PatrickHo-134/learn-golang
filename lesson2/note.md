## Lesson 2
This section covers Chapter 2 of Head First Go

### Content covered
- Conditional statements
- Block and variable scope
- Short varialbe declaration
- Package names (i.e. `fmt`) vs. import paths (i.e. `math/rand`)
- Loops


### Key notes:
- methods: functions that are associated with values of a given type. Go methods are kind of like the methods that you may have seen attached to “objects” in other languages, but they’re a bit simpler.
- Go doesn’t allow us to declare a variable unless we use it.
- Go treats everything from a `//` marker to the end of the line as a comment—and ignores it.
- Multiline comments start with `/*` and end with `*/`. Everything in between, including newlines, is ignored.
- It’s conventional to include a comment at the top of every program, explaining what it does.
- Unlike most programming languages, Go allows multiple return values from a function or method call.
- One common use of multiple return values is to return the function’s main result, and then a second value indicating whether there was an error.
- To discard a value without using it, use the _ blank identifier. The blank identifier can be used in place of any variable in any assignment statement.
- Avoid giving variables the same name as types, functions, or packages; it causes the variable to shadow (override) the item with the same name.
- Functions, conditionals, and loops all have blocks of code that appear within {} braces.
- Their code doesn’t appear within `{}` braces, but files and packages also comprise blocks.
- The scope of a variable is limited to the block it is defined within, and all blocks nested within that block.
- In addition to a name, a package may have an import path that is required when it is imported.
- The continue keyword skips to the next iteration of a loop.
- The break keyword exits out of a loop entirely.