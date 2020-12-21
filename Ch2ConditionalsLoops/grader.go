package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main()  {
	fmt.Println("Here we GO: Loops")
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(now)
	fmt.Println(year)

	broken := "G# r#cks!"
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

	//skip var
	//reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')

	if true && false {
		fmt.Println("I won't be printed!")
	} else if true && true {
		fmt.Println("But I will be")
	} else {
		fmt.Println(" And I'm default statement")
	}

	fileinfo, err := os.Stat("my.txt")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileinfo.Size())


	/*
		Grader
	*/
	fmt.Print("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	var status string
	if err != nil {
		log.Fatal(err)
	} else {
		if grade >= 60 {
			status = "passed"
		} else {
			status = "failed"
		}
	}
	fmt.Println(status)

}

