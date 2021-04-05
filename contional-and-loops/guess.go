package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
        var success bool
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)

        success = false
	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left")
		fmt.Println("Make a guess: ")
		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("I've chosen a random number between 1 and 100.")
		fmt.Println("Can you guess it?")

		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			fmt.Println("Good job! You guessed it!")
                        success = true
                        break
		}
	}

        if !success {
          fmt.Println("Oh...sorry! You could not guess it. It was", target)
        }
}
