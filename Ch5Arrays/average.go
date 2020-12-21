package main

import (
	"fmt"
	"hfgo/Ch5Arrays/datafile"
	"log"
)

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "si"}
	for i:=0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}
	for index, value := range notes {
		fmt.Println(index, value)
	}

	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	//numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("SUM: ", sum)
	fmt.Printf("AVG: %0.2f\n", sum/float64(len(numbers)))
}