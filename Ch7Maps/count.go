// count tallies the number of times each line occurs within a file
package main

import (
	"datafile"
	"fmt"
	"log"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	var board = make(map[string]int)
	for _, line := range lines {
		board[line]++
	}

	for name, votes := range board {
		fmt.Println(name, votes)
	}

	//var names[] string
	//var counts[] int
	//for _, line := range lines {
	//	matched := false
	//	for i, name := range names{
	//		if name == line{
	//			counts[i]++
	//			matched = true
	//		}
	//	}
	//	if matched == false {
	//		names = append(names, line)
	//		counts = append(counts, 1)
	//	}
	//}
	//
	//for i, name := range names{
	//	fmt.Printf("%s : %d\n", name, counts[i])
	//}
	printRanks()
}

func checkLetters() {
	data := []string{"a", "c", "e", "e", "a", "e"}
	c := make(map[string]int)

	for _, item := range data {
		c[item]++
	}
	letters := []string{"a", "b", "c", "d", "e"}

	for _, letter := range letters {
		count, ok := c[letter]
		if !ok {
			fmt.Printf("%s: not found\n", letter)
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}

	}

}

func printRanks() {
	ranks := map[string]int{
		"bronze": 3,
		"silver": 2,
		"gold":   1,
	}
	for medal, rank := range ranks {
		fmt.Printf("The %s medal's rank is %d\n", medal, rank)
	}

}
