package main

import (
	"fmt"
)

func m(){
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, number := range numbers {
		fmt.Println(i, number)
	}

	var letters = []string{"a","b","c"}
	for i, letter := range letters{
		fmt.Println(i, letter)
	}

	array := [5]string{"a", "b", "c", "d", "e"}
	sliceA := array[1:4]
	sliceB := array[2:5]
	fmt.Println(array)
	fmt.Println(sliceA)
	fmt.Println(sliceB)
	sliceA[2] = "CC"
	fmt.Println(array)
	fmt.Println(sliceA)
	fmt.Println(sliceB)


	myVarFunc(1, "a", "b", "c")

	res := inRange(-10, 2, 1,-19,22,0,14,2,1)
	fmt.Println(res)

	fmt.Println(sum(9, 7))
	fmt.Println(sum(4, 1, 2))
}

//Variadic function
func myVarFunc(a int, b ...string){
	fmt.Println(a)
	fmt.Println(b)
}

func inRange(min float64, max float64, numbers ...float64) []float64{
	var result []float64
	for _, number := range numbers {
		if number >= min && number <= max {
			result = append(result, number)
		}
	}
	return result
}

func sum(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}