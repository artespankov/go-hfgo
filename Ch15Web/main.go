package main

import "fmt"

func sayHi(){
	fmt.Println("Hi")
}

func divide(a int, b int) float64{
	return float64(a) / float64(b)
}

func multiply(a int, b int) float64 {
	return float64(a * b)
}

func doMath(passedFunction func(int, int) float64, a int, b int) float64{
	res := passedFunction(a, b)
	fmt.Println("Passed function result:", res)
	return res
}

func main(){
	var greeterFunction func()
	greeterFunction = sayHi
	greeterFunction()

	var mathFunction func(int, int) float64

	mathFunction = divide
	doMath(mathFunction, 100, 10)

	mathFunction = multiply
	doMath(mathFunction, 10, 17)
}
