## Lesson 2
This section covers Chapter 4 of Head First Go

### Content covered
- Creating packages
- Constants exported by packages
- Installing program executables with “go install”
- Setting GOPATH
- Downloading and installing packages with "go get"
- Reading package documentation with "go doc"
- Documenting your packages with doc comments


### Key notes:
- By default, the workspace directory is a directory named `go` within your user’s home directory.

- You can use another directory as your workspace by setting the `GOPATH` environment variable.

- Go uses three subdirectories within the workspace: the `bin` directory holds compiled executable programs, the `pkg` directory holds compiled package code, and the `src` directory holds Go source code.

- The names of the directories within the `src` directory are used to form a package’s import path. Names of nested directories are separated by `/` characters in the import path.

- The package’s name is determined by the `package` clauses at the top of the source code files within the package directory. Except for the `main` package, the package name should be the same as the name of the directory that contains it.

- Package names should be all lowercase, and ideally consist of a single word.

- A package’s functions can only be called from outside that package if they’re **exported**. A function is exported if its name begins with a capital letter.

- A **constant** is a name referring to a value that will never change.

- The go install command compiles a package’s code and stores it in the `pkg` directory for general packages, or the `bin` directory for executable programs.

- A common convention is to use the URL where a package is hosted as its import path. This allows the `go get` tool to find, download, and install packages given only their import path.

- The `go doc` tool displays documentation for packages. Doc comments within the code are included in go doc’s output.