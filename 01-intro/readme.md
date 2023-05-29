# Program Structure #
    - package declaration
    - import dependency packages
    - package level declarations
    - package init function
    - main function
    - othere functions

# Run a go program
```
    go run <filename.go>
```
# Create a build
```
    go build <filename.go>
```
# Create a build in a different name
```
    go build -o <build_name> <filename.go>
```
# To get the list of environment variables used by the 'go' tool
```
    go env
    go env <var_1> <var_2>
```

# To update the environment variables
```
    go env -w <var_1>=<value_1> <var_2>=<value_2>
```
# Environment variables for cross-compilation
- GOARCH
- GOOS
# To get the list of supported platforms for cross-compilation
```
    go tool dist list
```
# To cross compile
```
    GOOS=<os> GOARCH=<arch> go build <filename.go>
    ex:
    GOOS=windows GOARCH=386 go build 01-intro.go
```
# Data Types
- bool
- string
- int
- int8
- int16
- int32
- int64
- uint
- uint8
- uint16
- uint32
- uint64
- float32
- float64
- complex64 ( real [float32] + imaginary [float32] )
- complex126 ( real [float64] + imaginary [float64] )
- byte (alias for uint8)
- rune (alias for int32) (unicode code point)

# Variable Declarations
## Function Scope
    - Can use := syntax
    - Cannot have unused variables
## Package Scope
    - Cannot use := syntax
    - Can have unused variables

# Constants
    - Can have unused constants at function / package scope