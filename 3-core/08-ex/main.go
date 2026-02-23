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
