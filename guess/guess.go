//Guess challenges a player to guess a random number
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
	rand.Seed(time.Now().Unix())
	target := rand.Intn(100)

	fmt.Println("I've Chosen a random number between 1 and 100. \nCan you guess it?")

	fmt.Print("\nMake a Guess: ")
	number := 0

	for guesses := 1; guesses < 10; guesses++ {
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		number, err = strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}

		if number < target {
			fmt.Println("Oops. Your guess was LOW")
			fmt.Print("\nMake another Guess: ")
		} else if number > target {
			fmt.Println("Oops. Your guess was HIGH")
			fmt.Print("\nMake another Guess: ")
		} else {
			fmt.Println("Success!!!")
			break
		}
	}
}
