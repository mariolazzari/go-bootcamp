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

### Exercise: guess a number

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess a number between 1 and 50:")

	// guess a number between 0 and 50
	numToGuess := rand.Intn(50)
	fmt.Println("Random number is:", numToGuess)

	// read number
	reader := bufio.NewReader(os.Stdin)

	for {
		num, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// cast number
		num = strings.TrimSpace(num)
		iNum, err := strconv.Atoi(num)
		if err != nil {
			log.Fatal(err)
		}

		if iNum == numToGuess {
			fmt.Println("You guessed!")
			break
		} else if iNum < numToGuess {
			fmt.Println("Higher")
		} else {
			fmt.Println("Lower")
		}
	}

}
```

### Strings

```go
package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var pl = fmt.Println

func main() {
	// ----- STRINGS -----
	// Strings are arrays of bytes []byte
	// Escape Sequences : \n \t \" \\
	sV1 := "A word"

	// Replacer that can be used on multiple strings
	// to replace one string with another
	replacer := strings.NewReplacer("A", "Another")
	sV2 := replacer.Replace(sV1)
	pl(sV2)

	// Get length
	pl("Length : ", len(sV2))

	// Contains string
	pl("Contains Another :", strings.Contains(sV2, "Another"))

	// Get first index match
	pl("o index :", strings.Index(sV2, "o"))

	// Replace all matches with 0
	// If -1 was 2 it would replace the 1st 2 matches
	pl("Replace :", strings.Replace(sV2, "o", "0", -1))

	// Remove whitespace characters from beginning and end of string
	sV3 := "\nSome words\n"
	sV3 = strings.TrimSpace(sV3)

	// Split at delimiter
	pl("Split :", strings.Split("a-b-c-d", "-"))

	// Upper and lowercase string
	pl("Lower :", strings.ToLower(sV2))
	pl("Upper :", strings.ToUpper(sV2))

	// Prefix or suffix
	pl("Prefix :", strings.HasPrefix("tacocat", "taco"))
	pl("Suffix :", strings.HasSuffix("tacocat", "cat"))

	// ----- RUNES -----
	// In Go characters are called Runes
	// Runes are unicodes that represent characters
	rStr := "abcdefg"

	// Runes in string
	pl("Rune Count :", utf8.RuneCountInString(rStr))

	// Print runes in string
	for i, runeVal := range rStr {
		// Get index, Rune unicode and character
		fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	}
}
```

### Dates and times

```go
package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {

	// ----- TIME -----
	// Get day, month, year and time data
	// Get current time
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())

	// Set a location to get time
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Time in New York %s\n", now.In(loc))

	// Change location to Shanghai
	loc, _ = time.LoadLocation("Asia/Shanghai")
	fmt.Printf("Time in Shanghai %s\n", now.In(loc))

	// Get times using different time standards
	locEST, _ := time.LoadLocation("EST")
	locUTC, _ := time.LoadLocation("UTC")
	locMST, _ := time.LoadLocation("MST")
	fmt.Printf("EST : %s\n", now.In(locEST))
	fmt.Printf("UTC : %s\n", now.In(locUTC))
	fmt.Printf("MST : %s\n", now.In(locMST))

	// Calculate time since birthdate
	// Year, month, day, hour, minute, second
	// nanosecond and time zone
	birthDate := time.Date(1974, time.December,
		21, 11, 30, 10, 0, time.Local)

	// Get difference between past date and now
	diff := now.Sub(birthDate)

	// Difference in days
	fmt.Printf("Days Alive: %d days\n",
		int(diff.Hours()/24))

	// Hours
	fmt.Printf("Hours Alive: %d hours\n",
		int(diff.Hours()))
}
```

### Arrays

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// ----- ARRAYS -----
	// Collection of values with the same data type
	// and the size can't be changed
	// Default values are 0, 0.0, false or ""

	// Declare integer array with 5 elements
	var arr1 [5]int

	// Assign value to index
	arr1[0] = 1

	// Declare and initialize
	arr2 := [5]int{1, 2, 3, 4, 5}

	// Get by index
	pl("Index 0 :", arr2[0])

	// Length
	pl("Arr Length :", len(arr2))

	// Iterate with index
	for i := range len(arr2) {
		pl(arr2[i])
	}

	// Iterate with range
	for i, v := range arr2 {
		fmt.Printf("%d : %d\n", i, v)
	}

	// Multidimensional Array
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}

	// Print multidimensional array
	for i := range 2 {
		for j := range 2 {
			pl(arr3[i][j])
		}
	}

	// String into slice of runes
	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune Array : %d\n", v)
	}

	// Byte array to string
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string :", bStr)
}
```

### Slices

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// ----- SLICES -----
	// Slices are like arrays but they can grow
	// var name []dataType
	// Create a slice with make
	sl1 := make([]string, 6)

	// Assign values by index
	sl1[0] = "Society"
	sl1[1] = "of"
	sl1[2] = "the"
	sl1[3] = "Simulated"
	sl1[4] = "Universe"

	// Size of slice
	pl("Slice Size :", len(sl1))

	// Cycle with for
	for i := range sl1 {
		pl(sl1[i])
	}

	// Cycle with range
	for _, x := range sl1 {
		pl(x)
	}

	// Create a slice literal
	sl2 := []int{12, 21, 1974}
	pl(sl2)

	// A slice points at an array and you can create a slice
	// of an array (A slice is a view of an underlying array)
	// You can have multiple slices point to the same array
	sArr := [5]int{1, 2, 3, 4, 5}
	// Start at 0 index up to but not including the 2nd index
	sl3 := sArr[0:2]
	pl(sl3)

	// Get slice from beginning
	pl("1st 3 :", sArr[:3])

	// Get slice to the end
	pl("Last 3 :", sArr[2:])

	// If you change the array the slice also changes
	sArr[0] = 10
	pl("sl3 :", sl3)

	// Changing the slice also changes the array
	sl3[0] = 1
	pl("sArr :", sArr)

	// Append a value to a slice (Also overwrites array)
	sl3 = append(sl3, 12)
	pl("sl3 :", sl3)
	pl("sArr :", sArr)

	// Printing empty slices will return nils which show
	// as empty slices
	sl4 := make([]string, 6)
	pl("sl4 :", sl4)
	pl("sl4[0] :", sl4[0])
}
```

### Functions

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- FUNCTIONS -----
func sayHello() {
	pl("Hello")
}

// Returns sum of values
func getSum(x int, y int) int {
	return x + y
}

// Return multiple values
func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

// Return potential error
func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		// Define error message returned with dummy value
		// for ans
		return 0, fmt.Errorf("You can't divide by zero")
	} else {
		// If no error return nil
		return x / y, nil
	}
}

// Variadic function
func getSum2(nums ...int) int {
	sum := 0
	// nums gets converted into a slice which is
	// iterated by range (More on slices later)
	for _, num := range nums {
		sum += num
	}
	return sum
}

func getArraySum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func changeVal(f3 int) int {
	f3 += 1
	return f3
}

func main() {
	// ----- FUNCTIONS -----
	// func funcName(parameters) returnType {BODY}
	// If you only need a function in the current package
	// start with lowercase letter
	// Letters and numbers in camelcase
	sayHello()
	pl(getSum(5, 4))
	f1, f2 := getTwo(5)
	fmt.Printf("%d %d\n", f1, f2)

	// Function that can return an error
	ans, err := getQuotient(5, 0)
	if err == nil {
		pl("5/4 =", ans)
	} else {
		pl(err)
		// End program
		// log.Fatal(err)
	}

	// Function receives unknown number of parameters
	// Variadic Function
	pl("Unknown Sum :", getSum2(1, 2, 3, 4))

	// Pass an array to a function by value
	vArr := []int{1, 2, 3, 4}
	pl("Array Sum :", getArraySum(vArr))

	// Go passes the value to functions so it isn't changed
	// even if the same variable name is used
	f3 := 5
	pl("f3 before func :", f3)
	changeVal(f3)
	pl("f3 after func :", f3)
}
```

### Exercise: hangman

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

/*
+---+
    |
    |
    |
   ===

Secret Word : ______
Incorrect Guesses :

Guess a Letter : a

Sorry Your Dead! The word is ZOMBIE
Yes the Secret Word is ZOMBIE

Please Enter Only One Letter
Please Enter a Letter
Please Enter a Letter you Haven't Guessed
*/

var pl = fmt.Println

var hmArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		" 0   |\n" +
		"/|\\  |\n" +
		"/ \\  |\n" +
		"    ===\n",
}

var wordArr = [7]string{
	"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER",
	"ZODIAC", "ZOMBIE", "FLUFF",
}

// Stores the random word to be guessed
var randWord string

// Stores all letters guessed
var guessedLetters string

// Stores correct guesses
var correctLetters []string

// Letters guessed that aren't in the randWord
var wrongGuesses []string

func getRandWord() string {
	// Returns seconds as int
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	// Get random word from array
	randWord = wordArr[rand.Intn(7)]

	// Generate correctLetters array with correct size
	correctLetters = make([]string, len(randWord))
	return randWord
}

func showBoard() {
	// Prints current hangman state
	pl(hmArr[len(wrongGuesses)])

	fmt.Print("Secret Word : ")
	// Prints spaces for missing letter or the
	// letter
	// Cycle and if is a character print it
	// If not print an underscore
	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}

	fmt.Print("\nIncorrect Guesses : ")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	pl()
}

func getUserLetter() string {
	// Setup buffered reader that gets text from the keyboard
	reader := bufio.NewReader(os.Stdin)
	// The guess the user makes
	var guess string

	for true {
		fmt.Print("\nGuess a Letter : ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		// Convert to upper so it is easier to search
		guess = strings.ToUpper(guess)

		// Remove spaces
		guess = strings.TrimSpace(guess)

		// Create regular expression that verifies
		// if a single character is a letter
		var IsLetter = regexp.MustCompile(`^[a-zA-Z]$`).MatchString

		if len(guess) != 1 {
			fmt.Println("Please Enter Only One Letter")
		} else if !IsLetter(guess) {
			fmt.Println("Please Enter a Letter")
		} else if strings.Contains(guessedLetters, guess) {
			fmt.Println("Please Enter a Letter you Haven't Guessed")
		} else {
			return guess
		}
	}
	return guess
}

// Returns all indices matching the substring
func getAllIndexes(theStr, subStr string) (indices []int) {
	if (len(subStr) == 0) || (len(theStr) == 0) {
		return indices
	}

	// Offset will find 1st matching index and then
	// move forward to find the rest
	offset := 0
	for {
		i := strings.Index(theStr[offset:], subStr)
		// If no matches or no more matches
		// since we cycled to the end
		// return indices
		if i == -1 {
			// fmt.Println(indices)
			return indices
		}
		offset += i
		// Add index
		indices = append(indices, offset)
		// Jump ahead the length of the substring
		offset += len(subStr)
	}
}

// Updates correctLetters string to display guessed
// letters an _s
func updateCorrectLetters(letter string) {
	indexMatches := getAllIndexes(randWord, letter)

	// Assign all matching indexes the
	// supplied letter
	for _, v := range indexMatches {
		correctLetters[v] = letter
	}
}

// Checks if array contains an empty value
func sliceHasEmptys(theSlice []string) bool {
	for _, v := range theSlice {
		if len(v) == 0 {
			return true
		}
	}
	return false
}

func main() {
	pl(getRandWord())

	for true {
		// Display opening board
		showBoard()

		// Get user guess
		guess := getUserLetter()

		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)

			// Check if user guessed all
			// letters
			if sliceHasEmptys(correctLetters) {
				fmt.Println("More Letters to Guess")
			} else {
				fmt.Println("Yes the Secret Word is", randWord)
				break
			}

		} else {
			// They guessed a letter not in the secret
			// word
			// Append to string
			guessedLetters += guess
			// Append to slice
			wrongGuesses = append(wrongGuesses, guess)

			// Check if users ran out of guesses
			if len(wrongGuesses) >= 6 {
				fmt.Println("Sorry Your Dead! The word is", randWord)
				break
			}
		}
	}
}
```

### Pointers

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal2(myPtr *int) {
	*myPtr = 12
}

// Receives array by reference and doubles values
func dblArrVals(arr *[4]int) {
	for x := range 4 {
		arr[x] *= 2
	}
}

func getAverage(nums ...float64) float64 {
	var sum float64 = 0.0
	var numSize float64 = float64(len(nums))

	for _, val := range nums {
		sum += val
	}
	return (sum / numSize)
}

func main() {
	// ----- POINTERS -----

	// You can pass by reference with the &
	// (Address of Operator)
	// Print amount and address for amount in memory
	f4 := 10
	pl("f4 :", f4)
	pl("f4 Address :", &f4)

	// Store a pointer (Pointer to int)
	var f4Ptr *int = &f4
	pl("f4 Address :", f4Ptr)

	// Print value at pointer
	pl("f4 Value :", *f4Ptr)

	// Assign value using pointer
	*f4Ptr = 11
	pl("f4 Value :", *f4Ptr)

	// Change value in function
	pl("f4 before function :", f4)
	changeVal2(&f4)
	pl("f4 after function :", f4)

	// Pass an array by reference
	pArr := [4]int{1, 2, 3, 4}
	dblArrVals(&pArr)
	pl(pArr)

	// Passing a slice to a function works just
	// like when using variadic functions
	// Just add ... after the slice when passing
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAverage(iSlice...))
}
```

### File IO

```go
package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	// ----- FILE IO -----
	// We can create, write and read from files

	// Create a file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Says to close the file after program ends or when
	// there is a closing curly bracket
	defer f.Close()

	// Create list of primes
	iPrimeArr := []int{2, 3, 5, 7, 11}
	// Create string array from int array
	var sPrimeArr []string
	for _, i := range iPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}

	// Cycle through strings and write to file
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	// Open the created file
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read from file and print once per line
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime :", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}

	// Append to file
	/*
		 Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified
		O_RDONLY : open the file read-only
		O_WRONLY : open the file write-only
		O_RDWR   : open the file read-write
		These can be or'ed
		O_APPEND : append data to the file when writing
		O_CREATE : create a new file if none exists
		O_EXCL   : used with O_CREATE, file must not exist
		O_SYNC   : open for synchronous I/O
		O_TRUNC  : truncate regular writable file when opened
	*/

	// Check if file exists
	_, err = os.Stat("data.txt")
	if errors.Is(err, os.ErrNotExist) {
		pl("File Doesn't Exist")
	} else {
		f, err = os.OpenFile("data.txt",
			os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		if _, err := f.WriteString("13\n"); err != nil {
			log.Fatal(err)
		}
	}
}
```

### Command line

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(os.Args)
	// Get all values after the first index
	args := os.Args[1:]

	// Create int array from string array
	var iArgs = []int{}
	for _, i := range args {
		val, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		iArgs = append(iArgs, val)
	}

	max := 0
	for _, val := range iArgs {
		if val > max {
			max = val
		}
	}
	fmt.Println("Max Value :", max)
}
```

### Packages

```go
package main

import (
	"fmt"
	_ stuff "example/project/mypackage"
)

var pl = fmt.Println

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2,3,5,7,11}
	strArr := stuff.IntArrToStrArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.Typeof(strArr))
}
```

### Maps

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// ----- MAPS -----
	// Maps are collections of key/value pairs
	// Keys can be any data type that can be compared
	// using == (They can be a different type than
	// the value)
	// var myMap map [keyType]valueType

	// Declare a map variable
	var heroes map[string]string
	// Create the map
	heroes = make(map[string]string)

	// You can do it in one step
	villians := make(map[string]string)

	// Add keys and values
	heroes["Batman"] = "Bruce Wayne"
	heroes["Superman"] = "Clark Kent"
	heroes["The Flash"] = "Barry Allen"
	villians["Lex Luther"] = "Lex Luther"

	// Define with map literal
	superPets := map[int]string{1: "Krypto",
		2: "Bat Hound"}

	// Get value with key (Use %v with Printf)
	fmt.Printf("Batman is %v\n", heroes["Batman"])

	// If you access a key that doesn't exist
	// you get nil
	pl("Chip :", superPets[3])

	// You can check if there is a value or nil
	_, ok := superPets[3]
	pl("Is there a 3rd pet :", ok)

	// Cycle through map
	for k, v := range heroes {
		fmt.Printf("%s is %s\n", k, v)
	}

	// Delete a key value
	delete(heroes, "The Flash")
}
```

### Generics

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- FUNCTION THAT EXCEPTS GENERICS -----
// This generic type parameter is capital, between
// square brackets and has a rule for what data
// it will except called a constraint
// any : anything
// comparable : Anything that supports ==
// More Constraints : pkg.go.dev/golang.org/x/exp/constraints

// You can also define what is excepted like this
// Define that my generic must be an int or float64
type MyConstraint interface {
	int | float64
}

func getSumGen[T MyConstraint](x T, y T) T {
	return x + y
}

func main() {
	// ----- GENERICS -----
	// We can specify the data type to be used at a
	// later time with generics
	// It is mainly used when we want to create
	// functions that can work with
	// multiple data types
	pl("5 + 4 =", getSumGen(5, 4))
	pl("5.6 + 4.7 =", getSumGen(5.6, 4.7))

	// This causes an error
	// pl("5.6 + 4.7 =", getSumGen("5.6", "4.7"))
}
```

### Structs

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- STRUCTS -----
type customer struct {
	name    string
	address string
	bal     float64
}

// This struct has a function associated
type rectangle struct {
	length, height float64
}

func (r rectangle) Area() float64 {
	return r.length * r.height
}

// Customer passed as values
func getCustInfo(c customer) {
	fmt.Printf("%s owes us %.2f\n", c.name, c.bal)
}

func newCustAdd(c *customer, address string) {
	c.address = address
}

// Struct composition : Putting a struct in another
type contact struct {
	fName string
	lName string
	phone string
}

type business struct {
	name    string
	address string
	contact
}

func (b business) info() {
	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.lName)
}

func main() {
	// ----- STRUCTS -----
	// Structs allow you to store values with many
	// data types

	// Add values
	var tS customer
	tS.name = "Tom Smith"
	tS.address = "5 Main St"
	tS.bal = 234.56

	// Pass to function as values
	getCustInfo(tS)
	// or as reference
	newCustAdd(&tS, "123 South st")
	pl("Address :", tS.address)

	// Create a struct literal
	sS := customer{"Sally Smith", "123 Main", 0.0}
	pl("Name :", sS.name)

	// Structs with functions
	rect1 := rectangle{10.0, 15.0}
	pl("Rect Area :", rect1.Area())

	// Go doesn't support inheritance, but it does
	// support composition by embedding a struct
	// in another
	con1 := contact{
		"James",
		"Wang",
		"555-1212",
	}

	bus1 := business{
		"ABC Plumbing",
		"234 North St",
		con1,
	}

	bus1.info()
}
```

### Define types

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- DEFINED TYPES -----
// I'll define different cooking measurement types
// so we can do conversions
type Tsp float64
type TBs float64
type ML float64

// Convert with functions (Bad Way)
func tspToML(tsp Tsp) ML {
	return ML(tsp * 4.92)
}

func TBToML(tbs TBs) ML {
	return ML(tbs * 14.79)
}

// Associate method with types
func (tsp Tsp) ToMLs() ML {
	return ML(tsp * 4.92)
}
func (tbs TBs) ToMLs() ML {
	return ML(tbs * 14.79)
}

func main() {
	// ----- DEFINED TYPES -----
	// We used a defined type previously with structs
	// You can use them also to enhance the quality
	// of other data types
	// We'll create them for different measurements

	// Convert from tsp to mL
	ml1 := ML(Tsp(3) * 4.92)
	fmt.Printf("3 tsps = %.2f mL\n", ml1)

	// Convert from TBs to mL
	ml2 := ML(TBs(3) * 14.79)
	fmt.Printf("3 TBs = %.2f mL\n", ml2)

	// You can use arithmetic and comparison
	// operators
	pl("2 tsp + 4 tsp =", Tsp(2) + Tsp(4))
	pl("2 tsp > 4 tsp =", Tsp(2) > Tsp(4))

	// We can convert with functions
	// Bad Way
	fmt.Printf("3 tsp = %.2f mL\n", tspToML(3))
	fmt.Printf("3 TBs = %.2f mL\n", TBToML(3))

	// We can solve this by using methods which
	// are functions associated with a type
	tsp1 := Tsp(3)
	fmt.Printf("%.2f tsp = %.2f mL\n", tsp1, tsp1.ToMLs())
}
```

### Interfaces

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- INTERFACES -----
type Animal interface {
	AngrySound()
	HappySound()
}

// Define type with interface methods and its
// own method
type Cat string

func (c Cat) Attack() {
	pl("Cat Attacks its Prey")
}

// Return the cats name with a type conversion
func (c Cat) Name() string {
	return string(c)
}

func (c Cat) AngrySound() {
	pl("Cat says Hissssss")
}
func (c Cat) HappySound() {
	pl("Cat says Purrr")
}

func main() {
	// ----- INTERFACES -----
	// Interfaces allow you to create contracts
	// that say if anything inherits it that
	// they will implement defined methods

	// If we had animals and wanted to define that
	// they all perform certain actions, but in their
	// specific way we could use an interface

	// With Go you don't have to say a type uses
	// an interface. When your type implements
	// the required methods it is automatic
	var kitty Animal
	kitty = Cat("Kitty")
	kitty.AngrySound()

	// We can only call methods defined in the
	// interface for Cats because of the contract
	// unless you convert Cat back into a concrete
	// Cat type using a type assertion
	var kitty2 Cat = kitty.(Cat)
	kitty2.Attack()
	pl("Cats Name :", kitty2.Name())
}
```

### Concurrency

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var pl = fmt.Println

// ----- CONCURRENCY -----
func printTo15() {
	for i := 1; i <= 15; i++ {
		pl("Func 1 :", i)
	}
}
func printTo10() {
	for i := 1; i <= 10; i++ {
		pl("Func 2 :", i)
	}
}

// These functions will print in order using
// channels
// Func receives a channel and then sends values
// over channels once each time it is called
func nums1(channel chan int) {
	channel <- 1
	channel <- 2
	channel <- 3
}
func nums2(channel chan int) {
	channel <- 4
	channel <- 5
	channel <- 6
}

// ----- BANK ACCOUNT EXAMPLE -----
// Here I'll simulate customers accessing a
// bank account and lock out customers to
// allow for individual access
type Account struct {
	balance int
	lock    sync.Mutex // Mutual exclusion
}

func (a *Account) GetBalance() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.balance
}

func (a *Account) Withdraw(v int) {
	a.lock.Lock()
	defer a.lock.Unlock()
	if v > a.balance {
		pl("Not enough money in account")
	} else {
		fmt.Printf("%d withdrawn : Balance : %d\n",
			v, a.balance)
		a.balance -= v
	}
}

func main() {
	// ----- CONCURRENCY -----
	// Concurrency allows us to have multiple
	// blocks of code share execution time by
	// pausing their execution. We can also
	// run blocks of codes in parallel at the same
	// time. (Concurrent tasks in Go are called
	// goroutines)

	// To execute multiple functions in new
	// goroutines add the word go in front of
	// the function calls (Those functions can't
	// have return values)

	// We can't control when functions execute
	// so we may get different results
	go printTo15()
	go printTo10()

	// We have to pause the main function because
	// if main ends so will the goroutines
	time.Sleep(2 * time.Second) // Pause 2 seconds

	// You can have goroutines communicate by
	// using channels. The sending goroutine
	// also makes sure the receiving goroutine
	// receives the value before it attempts
	// to use it

	// Create a channel : Only carries values of
	// 1 type
	channel1 := make(chan int)
	channel2 := make(chan int)
	go nums1(channel1)
	go nums2(channel2)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel1)
	pl(<-channel2)
	pl(<-channel2)
	pl(<-channel2)

	// Using locks to protect data from being
	// accessed by more than one user at a time
	// Locks are another option when you don't
	// have to pass data
	var acct Account
	acct.balance = 100
	pl("Balance :", acct.GetBalance())

	for range 12 {
		go acct.Withdraw(10)
	}
	time.Sleep(2 * time.Second)
}
```

### Closures

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- CLOSURES -----
// Pass a function to a function
func useFunc(f func(int, int) int, x, y int) {
	pl("Answer :", (f(x, y)))
}

func sumVals(x, y int) int {
	return x + y
}

func main() {
	// ----- CLOSURES -----
	// Closures are functions that don't have to be
	// associated with an identifier (Anonymous)

	// Create a closure that sums values
	intSum := func(x, y int) int { return x + y }
	pl("5 + 4 =", intSum(5, 4))

	// Closures can change values outside the function
	samp1 := 1
	changeVar := func() { samp1 += 1 }
	changeVar()
	pl("samp1 =", samp1)

	// Pass a function to a function
	useFunc(sumVals, 5, 8)
}
```

### Recursion

```go
package main

import (
	"fmt"
)

var pl = fmt.Println

// ----- RECURSION -----
func factorial(num uint64) uint64 {
	// This condition ends calling functions
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	// ----- RECURSION -----
	// Recursion occurs when a function calls itself
	// There must be a condition that ends this
	// Finding a factorial is commonly used
	pl("Factorial 4 =", factorial(4))
	// 1st : result = 4 * factorial(3) = 4 * 6 = 24
	// 2nd : result = 3 * factorial(2) = 3 * 2 = 6
	// 3rd : result = 2 * factorial(1) = 2 * 1 = 2
}
```

### Regular expressions

```go
package main

import (
	"fmt"
	"regexp"
)

var pl = fmt.Println

func main() {
	// ----- REGULAR EXPRESSIONS -----
	// You can use regular expressions to test
	// if a string matches a pattern

	// Search for ape followed by not a space
	reStr := "The ape was at the apex"
	match, _ := regexp.MatchString("(ape[^ ]?)", reStr)
	pl(match)

	// You can compile them
	// Find multiple words ending with at
	reStr2 := "Cat rat mat fat pat"
	r, _ := regexp.Compile("([crmfp]at)")

	// Did you find any matches?
	pl("MatchString :", r.MatchString(reStr2))

	// Return first match
	pl("FindString :", r.FindString(reStr2))

	// Starting and ending index for 1st match
	pl("Index :", r.FindStringIndex(reStr2))

	// Return all matches
	pl("All String :", r.FindAllString(reStr2, -1))

	// Get 1st 2 matches
	pl("All String :", r.FindAllString(reStr2, 2))

	// Get indexes for all matches
	pl("All Submatch Index :", r.FindAllStringSubmatchIndex(reStr2, -1))

	// Replace all matches with Dog
	pl(r.ReplaceAllString(reStr2, "Dog"))
}
```

### Automated testing

```go
package main

import (
	"testing"
)

func TestIsEmail(t *testing.T) {
	_, err := IsEmail("hello")
	if err == nil {
		t.Error("hello is not an email")
	}

	_, err = IsEmail("derek@aol.com")
	if err != nil {
		t.Error("derek@aol.com is an email")
	}

	_, err = IsEmail("derek@aol")
	if err != nil {
		t.Error("derek@aol is not email")
	}
}
```

## Sudoku

### Version 1

```go
package main

import (
	"fmt"
	"strconv"
)

/*
SUDOKU RULES
1. 9 by 9 square
2. Each row and column must contain numbers 1-9
3. Each 3x3 square must contain numbers 1-9
4. No repeats allowed in rows, columns or squares
*/

// Outputs the current board to screen
func displayBoard(puzz [][]int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := range puzz {
		// Cycle through columns in puzz
		for j := range puzz_cols {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Print("\n")
	}
}

func getEmptySpace(puzz [][]int) (int, int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := range puzz {
		// Cycle through columns in puzz
		for j := range puzz_cols {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Receives slice, number to check, row &
// column and decides if that number fits
// the rules of the game
func isNumValid(puzz [][]int, guess int, row int, column int) bool {

	// Cycle through all values in row to see
	// if the row is valid (We only want index
	// which is the 1st value)
	for index := range puzz {
		// Use index for cycling
		// Check if any of the values are equal
		// to our guess and also that we didn't
		// already place it there in this
		// function
		if puzz[row][index] == guess && column != index {
			return false
		}
	}
	return true
}

func main() {
	// Holds starting puzzle
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	// ----- TESTING FUNCTIONS -----
	// Displays current board
	displayBoard(puzz)
	// Returns first open space
	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Puzzle is Solved")
	}
	// Testing if a guess is valid
	// Is 1 valid for row 1
	fmt.Println(isNumValid(puzz, 1, 0, 0))
	// Is 7 valid for row 1
	fmt.Println(isNumValid(puzz, 7, 0, 0))
}
```

### Version 2

```go
package main

import (
	"fmt"
	"strconv"
)

/*
SUDOKU RULES
1. 9 by 9 square
2. Each row and column must contain numbers 1-9
3. Each 3x3 square must contain numbers 1-9
4. No repeats allowed in rows, columns or squares
*/

// Outputs the current board to screen
func displayBoard(puzz [][]int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := range puzz {
		// Cycle through columns in puzz
		for j := range puzz_cols {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Print("\n")
	}
}

func getEmptySpace(puzz [][]int) (int, int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := range puzz {
		// Cycle through columns in puzz
		for j := range puzz_cols {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Receives slice, number to check, row &
// column and decides if that number fits
// the rules of the game
func isNumValid(puzz [][]int, guess int, row int, column int) bool {

	// Cycle through all values in row to see
	// if the row is valid (We only want index
	// which is the 1st value)
	for index := range puzz {
		// Use index for cycling
		// Check if any of the values are equal
		// to our guess and also that we didn't
		// already place it there in this
		// function
		if puzz[row][index] == guess && column != index {
			return false
		}
	}

	for index := range puzz {
		if puzz[index][column] == guess && column != index {
			return false
		}
	}

	// Is number valid for box?
	// Row 0 & Column 3
	// Row - (Row % 3) + Value for cycling (0-2)
	// 0 - (0 % 3) + 0 = 0 - 0 + 0 = 0 (1st row in box)
	// 0 - (0 % 3) + 1 = 0 - 0 + 1 = 1 (1st row in box)
	// 0 - (0 % 3) + 2 = 0 - 0 + 2 = 2 (1st row in box)

	// Col - (Col % 3) + Value for cycling (0-2)
	// 3 - (3 % 3) + 0 = 3 - 0 + 0 = 3 (1st col in box)
	// 3 - (3 % 3) + 1 = 3 - 0 + 1 = 4 (1st col in box)
	// 3 - (3 % 3) + 2 = 3 - 0 + 2 = 5 (1st col in box)

	for k := range 3 {
		for l := range 3 {
			if puzz[row-row%3+k][column-column%3+l] == guess && (row-row%3+k != row || column-column%3+l != column) {
				return false
			}
		}
	}

	return true
}

func main() {

	// Holds starting puzzle
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	// ----- TESTING FUNCTIONS -----
	// Displays current board
	displayBoard(puzz)
	// Returns first open space
	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Puzzle is Solved")
	}
	// Testing if a guess is valid
	fmt.Println(isNumValid(puzz, 1, 0, 6))

}
```

### Version 3

```go
package main

import (
	"fmt"
	"strconv"
)

/*
SUDOKU RULES
1. 9 by 9 square
2. Each row and column must contain numbers 1-9
3. Each 3x3 square must contain numbers 1-9
4. No repeats allowed in rows, columns or squares
*/

// Outputs the current board to screen
func displayBoard(puzz [][]int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := range puzz {
		// Cycle through columns in puzz
		for j := range puzz_cols {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Print("\n")
	}
}

func getEmptySpace(puzz [][]int) (int, int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := range puzz {
		// Cycle through columns in puzz
		for j := range puzz_cols {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Receives slice, number to check, row &
// column and decides if that number fits
// the rules of the game
func isNumValid(puzz [][]int, guess int, row int, column int) bool {

	// Cycle through all values in row to see
	// if the row is valid (We only want index
	// which is the 1st value)
	for index := range puzz {
		// Use index for cycling
		// Check if any of the values are equal
		// to our guess and also that we didn't
		// already place it there in this
		// function
		if puzz[row][index] == guess && column != index {
			return false
		}
	}

	for index := range puzz {
		if puzz[index][column] == guess && column != index {
			return false
		}
	}

	// Is number valid for box?
	// Row 0 & Column 3
	// Row - (Row % 3) + Value for cycling (0-2)
	// 0 - (0 % 3) + 0 = 0 - 0 + 0 = 0 (1st row in box)
	// 0 - (0 % 3) + 1 = 0 - 0 + 1 = 1 (1st row in box)
	// 0 - (0 % 3) + 2 = 0 - 0 + 2 = 2 (1st row in box)

	// Col - (Col % 3) + Value for cycling (0-2)
	// 3 - (3 % 3) + 0 = 3 - 0 + 0 = 3 (1st col in box)
	// 3 - (3 % 3) + 1 = 3 - 0 + 1 = 4 (1st col in box)
	// 3 - (3 % 3) + 2 = 3 - 0 + 2 = 5 (1st col in box)

	for k := range 3 {
		for l := range 3 {
			if puzz[row-row%3+k][column-column%3+l] == guess && (row-row%3+k != row || column-column%3+l != column) {
				return false
			}
		}
	}
	return true
}

/*
-------------------------
| 0 0 0 | 0 3 5 | 0 7 0 |
| 2 5 0 | 0 4 6 | 8 0 1 |
| 0 1 3 | 7 0 8 | 0 4 9 |
-------------------------
| 1 9 0 | 0 0 7 | 0 0 4 |
| 0 0 5 | 0 0 2 | 0 9 6 |
| 8 0 2 | 0 9 4 | 0 0 7 |
-------------------------
| 3 7 0 | 0 0 9 | 0 0 0 |
| 0 6 1 | 0 7 0 | 0 0 0 |
| 4 0 0 | 5 8 1 | 0 0 0 |
-------------------------

// Recursion Problem (Solution)
// 1. Cycle across the rows column by column (1-9)
// 2. Check if valid number
// a. If true update array
// b. If false change back to 0
// c. If false find a new value for previous we
// thought was correct
// 3. Check next column
*/

func main() {

	// Holds starting puzzle
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	// ----- TESTING FUNCTIONS -----
	// Displays current board
	displayBoard(puzz)
	// Returns first open space
	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Puzzle is Solved")
	}
	// Testing if a guess is valid
	fmt.Println(isNumValid(puzz, 1, 0, 6))

}
```

### Version 4: final

```go
package main

import (
	"fmt"
	"strconv"
)

/*
SUDOKU RULES
1. 9 by 9 square
2. Each row and column must contain numbers 1-9
3. Each 3x3 square must contain numbers 1-9
4. No repeats allowed in rows, columns or squares
*/

// Outputs the current board to screen
func displayBoard(puzz [][]int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := range puzz {
		// Cycle through columns in puzz
		for j := range puzz_cols {
			fmt.Print(strconv.Itoa(puzz[i][j]) + " ")
		}
		fmt.Print("\n")
	}
}

func getEmptySpace(puzz [][]int) (int, int) {
	puzz_cols := len(puzz[0])
	// Cycles through rows in the puzz
	for i := range puzz {
		// Cycle through columns in puzz
		for j := range puzz_cols {
			if puzz[i][j] == 0 {
				return i, j
			}
		}
	}
	return -1, -1
}

// Receives slice, number to check, row &
// column and decides if that number fits
// the rules of the game
func isNumValid(puzz [][]int, guess int, row int, column int) bool {

	// Cycle through all values in row to see
	// if the row is valid (We only want index
	// which is the 1st value)
	for index := range puzz {
		// Use index for cycling
		// Check if any of the values are equal
		// to our guess and also that we didn't
		// already place it there in this
		// function
		if puzz[row][index] == guess && column != index {
			return false
		}
	}

	// NEW Check if column is valid
	for index := range puzz {
		if puzz[index][column] == guess && column != index {
			return false
		}
	}

	// NEW Check if the box is valid
	// We need to convert all rows and columns
	// into their 3x3 version so we can cycle
	// through them

	// We need to do math to generate the row and column
	// numbers to check in our slice

	// Let's use Row 0 Column 3 as an example
	// Row - (Row % 3) + Val for Cycling (0-2)
	// 0 - (0 % 3) + 0 = 0 - 0 + 0 = 0 (1st row in box)
	// 0 - (0 % 3) + 1 = 0 - 0 + 1 = 1 (2nd row in box)
	// 0 - (0 % 3) + 2 = 0 - 0 + 2 = 2 (3rd row in box)

	// Col - (Col % 3) + Val for Cycling (0-2)
	// 3 - (3 % 3) + 0 = 3 - 0 + 0 = 3 (1st col in box)
	// 3 - (3 % 3) + 1 = 3 - 0 + 1 = 4 (2nd col in box)
	// 3 - (3 % 3) + 2 = 3 - 0 + 2 = 5 (3rd col in box)

	for k := range 3 {
		for l := range 3 {
			// We need to check if any element in the box
			// is equal to what we just added but skip
			// checking the position we just placed our
			// value in
			if puzz[row-row%3+k][column-column%3+l] == guess && (row-row%3+k != row || column-column%3+l != column) {
				return false
			}
		}
	}

	return true
}

/*
-------------------------
| 0 0 0 | 0 3 5 | 0 7 0 |
| 2 5 0 | 0 4 6 | 8 0 1 |
| 0 1 3 | 7 0 8 | 0 4 9 |
-------------------------
| 1 9 0 | 0 0 7 | 0 0 4 |
| 0 0 5 | 0 0 2 | 0 9 6 |
| 8 0 2 | 0 9 4 | 0 0 7 |
-------------------------
| 3 7 0 | 0 0 9 | 0 0 0 |
| 0 6 1 | 0 7 0 | 0 0 0 |
| 4 0 0 | 5 8 1 | 0 0 0 |
-------------------------
6 4 8 1 2 #9 (Change 9 to 0 and try another 2)
9 2 (1st row done)
7 9 3 (2nd row done)
#9
-------------------------
| 6 4 8 | 1 3 5 | 9 7 2 |
| 2 5 7 | 9 4 6 | 8 3 1 |
| 0 1 3 | 7 0 8 | 0 4 9 |
-------------------------
*/

/*
9 8 4 1 3 5 6 7 2
2 5 7 9 4 6 8 3 1
6 1 3 7 2 8 5 4 9
1 9 6 3 5 7 2 8 4
7 4 5 8 1 2 3 9 6
8 3 2 6 9 4 1 5 7
3 7 8 2 6 9 4 1 5
5 6 1 3 7 4 9 2 8
4 2 9 5 8 1 7 6 3
*/

// Recursively solves puzzle by placing values
// if it reaches a point where there is no
// valid answer it changes the previous
// supposedly correct answer back and tries
// another value for what we previously though was
// correct

// By using this backtracking when we get to the 9 9
// position on our board we know we solved the puzzle
func solvePuzzle(puzz [][]int) bool {
	// Check for empty space and if there
	// is none return true meaning the puzzle
	// is solved, otherwise solve the puzzle
	row, column := getEmptySpace(puzz)
	if row == -1 {
		return true
	} else {
		row, column = getEmptySpace(puzz)
	}

	// Check if we can add a 1 through 9 into the
	// empty space
	for i := 1; i <= 9; i++ {
		// Is this a valid number?
		if isNumValid(puzz, i, row, column) {
			// If so put it in the empty space
			puzz[row][column] = i

			// Continue until solvePuzzle with new puzz
			// until it returns true meaning we no longer
			// have empty spaces and we are finished
			if solvePuzzle(puzz) {
				return true
			}
			// If solvePuzzle doesn't return true we
			// reset to an empty space and backtrack
			// and try the next value in the for loop
			puzz[row][column] = 0
		}
	}
	return false
}

func main() {

	// Holds starting puzzle
	puzz := [][]int{
		{0, 0, 0, 0, 3, 5, 0, 7, 0},
		{2, 5, 0, 0, 4, 6, 8, 0, 1},
		{0, 1, 3, 7, 0, 8, 0, 4, 9},
		{1, 9, 0, 0, 0, 7, 0, 0, 4},
		{0, 0, 5, 0, 0, 2, 0, 9, 6},
		{8, 0, 2, 0, 9, 4, 0, 0, 7},
		{3, 7, 0, 0, 0, 9, 0, 0, 0},
		{0, 6, 1, 0, 7, 0, 0, 0, 0},
		{4, 0, 0, 5, 8, 1, 0, 0, 0},
	}

	// ----- TESTING FUNCTIONS -----
	// Displays current board
	displayBoard(puzz)
	// Returns first open space
	row, _ := getEmptySpace(puzz)
	if row != -1 {
		fmt.Println(getEmptySpace(puzz))
	} else {
		fmt.Println("Puzzle is Solved")
	}
	// Testing if a guess is valid
	// Is 1 valid for row 1
	// fmt.Println(isNumValid(puzz, 1, 0, 0))
	// Is 7 valid for row 1
	// fmt.Println(isNumValid(puzz, 7, 0, 0))

	// NEW
	// Testing if row and column is valid
	fmt.Println(isNumValid(puzz, 1, 0, 0))
	fmt.Println(isNumValid(puzz, 6, 0, 0))
	fmt.Println(isNumValid(puzz, 9, 0, 0))
	fmt.Println(isNumValid(puzz, 2, 0, 1))

	fmt.Println(isNumValid(puzz, 7, 4, 0))
	fmt.Println(isNumValid(puzz, 8, 4, 1))

	displayBoard(puzz)
	fmt.Println()
	solvePuzzle(puzz)
	displayBoard(puzz)
}
```

## Web apps

### Web app 1

```sh
go mod init github.com/mariolazzari/go-bootcamp/webapp1
```

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, err := fmt.Fprintf(w, "Ciao Mario!")
		if err != nil {
			log.Fatalf("Error writing response", err)
		}
		log.Println("Bytes writte:", bytes)
	})

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatalln("Error starting server:", err)
	}
}
```

### Web App 2

```go
package main

import (
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func write(writer http.ResponseWriter,
	msg string) {
	_, err := writer.Write([]byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, err := template.ParseFiles("./templates/" + tmpl)
	errorCheck(err)
	err = parsedTemplate.Execute(w, nil)
	errorCheck(err)
}

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func homeHandler(w http.ResponseWriter,
	request *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func addHandler(writer http.ResponseWriter,
	request *http.Request) {
	write(writer, "Hello internet\n")
	sum := getSum(5, 4)
	output := fmt.Sprintf("5 + 4 = %d\n", sum)
	write(writer, output)
}

func getSum(x, y int) int {
	return x + y
}

func divideHandler(writer http.ResponseWriter,
	request *http.Request) {
	v, err := getQuotient(5, 4)
	if err != nil {
		write(writer, "can't divide by zero\n")
	} else {
		output := fmt.Sprintf("5 / 4 = %.2f\n", v)
		write(writer, output)
	}
}

func getQuotient(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	} else {
		return (x / y), nil
	}
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/getsum", addHandler)
	http.HandleFunc("/divide", divideHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
```

## Complex web app

```sh
go mod init github.com/mariolazzari/go-bootcamp/webapp3
```
