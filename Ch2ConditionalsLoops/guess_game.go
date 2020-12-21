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

// guess challenges players to guess a random number
func main()  {
	seconds := time.Now().Unix();
	rand.Seed(seconds)
	goal := rand.Intn(100) + 1

	reader := bufio.NewReader(os.Stdin)
	success := false
	for x:=10; x>=1; x-- {
		fmt.Println("Turns left: ", x)
		fmt.Print("Your guess?: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(err)
		}
		if guess < goal {
			println("Your guess was LOW")
		} else if guess > goal {
			println("Your guess was HIGH")
		} else {
			println("You've WON")
			success = true
			break
		}
	}
	if !success {
		fmt.Println("You lost. The goal is: ", goal)
	}

}
