# Go programming bootcamp

## Core Go language

### Introduction

```go
// A package is a collection of code
// We can define what package we want our code to belong to
// We use main when we want our code to run in the terminal
package main

// Import multiple packages
// You could use an alias like f "fmt"
import (
	"bufio"
	f "fmt"
	"log"
	"os"
)

// Create alias to long function names
var pl = f.Println

/*
I'm a block comment
*/

// When a Go program executes it executes a function named main
// Go statements don't require semicolons
func main() {
	// Prints text and a newline
	// List package name followed by a period and the function name
	pl("Hello Go")

	// Get user input (To run this in the terminal go run hellogo.go)
	pl("What is your name?")
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)
	// Copy text up to the newline
	// The blank identifier _ will get err and ignore it (Bad Practice)
	// name, _ := reader.ReadString('\n')
	// It is better to handle it
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		// Log this error
		log.Fatal(err)
	}
}
```

### Variables and types

```go
// A package is a collection of code
// We can define what package we want our code to belong to
// We use main when we want our code to run in the terminal
package main

// Import multiple packages
// You could use an alias like f "fmt"
import (
	"fmt"
	"reflect"
	"strconv"
)

// Create alias to long function names
var pl = fmt.Println

func main() {

	// ----- VARIABLES -----
	// var name type
	// Name must begin with letter and then letters or numbers
	// If a variable, function or type starts with a capital letter
	// it is considered exported and can be accessed outside the
	// package and otherwise is available only in the current package
	// Camal case is the default naming convention

	// var vName string = "Derek"
	// var v1, v2 = 1.2, 3.4

	// Short variable declaration (Type defined by data)
	// var v3 = "Hello"

	// Variables are mutable by default (Value can change as long
	// as the data type is the same)
	// v1 := 2.4

	// After declaring variables to assign values to them always use
	// = there after. If you use := you'll create a new variable

	// ----- DATA TYPES -----
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('🦍'))

	// ----- CASTING -----
	// To cast type the type to convert to with the variable to
	// convert in parentheses
	// Doesn't work with bools or strings
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	// Convert string to int (ASCII to Integer)
	// Returns the result with an error if any
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	// Convert int to string (Integer to ASCII)
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	// Convert string to float
	cV7 := "3.14"
	// Handling potential errors (Prints if err == nil)
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	// Use Sprintf to convert from float to string
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
}
```

### Formatting

```go
// A package is a collection of code
// We can define what package we want our code to belong to
// We use main when we want our code to run in the terminal
package main

// Import multiple packages
// You could use an alias like f "fmt"
import (
	"fmt"
	"reflect"
	"strconv"
)

// Create alias to long function names
var pl = fmt.Println

func main() {

	// ----- VARIABLES -----
	// var name type
	// Name must begin with letter and then letters or numbers
	// If a variable, function or type starts with a capital letter
	// it is considered exported and can be accessed outside the
	// package and otherwise is available only in the current package
	// Camal case is the default naming convention

	// var vName string = "Derek"
	// var v1, v2 = 1.2, 3.4

	// Short variable declaration (Type defined by data)
	// var v3 = "Hello"

	// Variables are mutable by default (Value can change as long
	// as the data type is the same)
	// v1 := 2.4

	// After declaring variables to assign values to them always use
	// = there after. If you use := you'll create a new variable

	// ----- DATA TYPES -----
	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""
	pl(reflect.TypeOf(25))
	pl(reflect.TypeOf(3.14))
	pl(reflect.TypeOf(true))
	pl(reflect.TypeOf("Hello"))
	pl(reflect.TypeOf('🦍'))

	// ----- CASTING -----
	// To cast type the type to convert to with the variable to
	// convert in parentheses
	// Doesn't work with bools or strings
	cV1 := 1.5
	cV2 := int(cV1)
	pl(cV2)

	// Convert string to int (ASCII to Integer)
	// Returns the result with an error if any
	cV3 := "50000000"
	cV4, err := strconv.Atoi(cV3)
	pl(cV4, err, reflect.TypeOf(cV4))

	// Convert int to string (Integer to ASCII)
	cV5 := 50000000
	cV6 := strconv.Itoa(cV5)
	pl(cV6)

	// Convert string to float
	cV7 := "3.14"
	// Handling potential errors (Prints if err == nil)
	if cV8, err := strconv.ParseFloat(cV7, 64); err == nil {
		pl(cV8)
	}

	// Use Sprintf to convert from float to string
	cV9 := fmt.Sprintf("%f", 3.14)
	pl(cV9)
}
```
