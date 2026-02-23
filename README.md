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

### Exercise: age

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// get customer age
	fmt.Println("Your age:")
	reader := bufio.NewReader(os.Stdin)

	strAge, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Error reading age:%s", err)
	}

	// trim spaces
	strAge = strings.TrimSpace(strAge)

	// check age and return school
	age, err := strconv.Atoi(strAge)
	if err != nil {
		log.Fatalf("Error converting age: %s", err)
	}

	if age < 5 {
		fmt.Println("You are too young for school")
	} else if age == 5 {
		fmt.Println("Go to kindergarten")
	} else if age > 5 && age <= 17 {
		grade := age - 5
		fmt.Printf("Go to %d grade\n", grade)
	} else {
		fmt.Println("Go to college")
	}
}
```

### Math

```go
package main

import (
	"fmt"
	"math"
	"math/rand"
)

var pl = fmt.Println

func main() {
	pl("5 + 4 =", 5+4)
	pl("5 - 4 =", 5-4)
	pl("5 * 4 =", 5*4)
	pl("5 / 4 =", 5/4)
	pl("5 % 4 =", 5%4)

	// Shorthand increment
	// Instead of mInt = mInt + 1 (mInt += 1)
	// -= *= /=
	mInt := 1
	mInt += 1
	// Also increment or decrement with ++ and --
	mInt++

	// Float precision increases with the size of your values
	pl("Float Precision =", 0.11111111111111111111+
		0.11111111111111111111)

	// Create a random value between 0 and 50
	// Get a seed value for our random number generator based on
	// seconds since 1/1/70 to make our random number more random
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)

	// There are many math functions
	pl("Abs(-10) =", math.Abs(-10))
	pl("Pow(4, 2) =", math.Pow(4, 2))
	pl("Sqrt(16) =", math.Sqrt(16))
	pl("Cbrt(8) =", math.Cbrt(8))
	pl("Ceil(4.4) =", math.Ceil(4.4))
	pl("Floor(4.4) =", math.Floor(4.4))
	pl("Round(4.4) =", math.Round(4.4))
	pl("Log2(8) =", math.Log2(8))
	pl("Log10(100) =", math.Log10(100))
	// Get the log of e to the power of 2
	pl("Log(7.389) =", math.Log(math.Exp(2)))
	pl("Max(5,4) =", math.Max(5, 4))
	pl("Min(5,4) =", math.Min(5, 4))

	// Convert 90 degrees to radians
	r90 := 90 * math.Pi / 180
	// Convert 1.5708 radians to degrees
	d90 := r90 * (180 / math.Pi)
	fmt.Printf("%f radians = %f degrees\n", r90, d90)
	pl("Sin(90) =", math.Sin(r90))

	// There are also functions for Cos, Tan, Acos, Asin
	// Atan, Asinh, Acosh, Atanh, Atan2, Cosh, Sinh, Sincos
	// Htpot
}
```

### Exercise: sum

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read from standard input
	reader := bufio.NewReader(os.Stdin)

	// read first number
	fmt.Println("Enter first number:")
	num1, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// trim spaces and convert to int
	num1 = strings.TrimSpace(num1)
	iNum1, err := strconv.Atoi(num1)
	if err != nil {
		log.Fatal(err)
	}

	// repeat for second number
	fmt.Println("Enter second number:")
	num2, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	num2 = strings.TrimSpace(num2)
	iNum2, err := strconv.Atoi(num2)
	if err != nil {
		log.Fatal(err)
	}

	sum := iNum1 + iNum2
	fmt.Printf("Sum: %d\n", sum)

}
```

### Looping

```go
package main

import "fmt"

var pl = fmt.Println

func main() {
	// ----- FOR LOOPS -----
	// for initialization; condition; postStatement {BODY}
	// Print numbers 1 through 5
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	// Do the opposite
	for x := 5; x >= 1; x-- {
		pl(x)
	}

	// x is out of the scope of the for loop so it doesn't exist
	// pl("x :", x)

	// For is used to create while loops as well
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	// Cycle through an array with range
	// More on arrays later
	// We don't need the index so we ignore it
	// with the blank identifier _
	aNums := []int{1, 2, 3}
	for _, num := range aNums {
		pl(num)
	}

	// We can allow a condition in the for loop
	// to decide when to exit
	xVal := 1
	for true {
		if xVal == 5 {
			// Break exits the loop
			break
		}
		pl(xVal)
		xVal++
	}
}
```

### Exercise: loop

```go

```
