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
