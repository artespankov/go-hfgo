package main

import (
	"fmt"
	"math"
)

//func main(){
//	var total float64
//	//total += paintNeeded(12, 32)
//	//total += paintNeeded(22, 11)
//	//total += paintNeeded(15, 17)
//	area, err := paintNeeded(10, 199)
//	if err != nil {
//		log.Fatal(err)
//	}
//	total += area
//	fmt.Printf("%.2f liters needed\n", area)
//	fmt.Printf("%.2f Total liters needed\n", total)
//
//	root, err := squareRoot(-9.3)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Printf("%.2f\n", root)
//	}
//}

func paintNeeded(width float64, height float64) (float64, error) {

	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / 10.0, nil
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("can't get square root of negative number")
	}
	return math.Sqrt(number), nil
}