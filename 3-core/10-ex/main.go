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
