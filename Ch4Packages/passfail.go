package main

import (
	"fmt"
	"hfgo/Ch4 - Packages/github.com/artespankov/keyboard"
	"log"
)

// reports whether a grade is passing or failing
func main() {
	fmt.Println("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}
	fmt.Println("A grade of ", grade, " is ", status)
}
