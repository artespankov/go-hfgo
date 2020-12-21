package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	avg(os.Args[1:]...)
}

func avg(args ...string){
	var sum float64 = 0
	for _, arg := range args{
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}
	fmt.Println("SUM: ", sum)
	fmt.Printf("AVG: %0.2f\n", sum/float64(len(args)))
}