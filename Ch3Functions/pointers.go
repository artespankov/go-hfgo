package main

import (
	"fmt"
)

func main(){
	//var myInt int
	//fmt.Println(&myInt)
	//fmt.Println(*&myInt)
	//var myStr string
	//fmt.Println(&myStr)
	//fmt.Println(*&myStr)
	//var myBool bool
	//fmt.Println(&myBool)
	//fmt.Println(*&myBool)

	// Update value at a pointer
	//myFloat := 4.2
	//fmt.Println(myFloat)
	//myFloatPointer := &myFloat
	//*myFloatPointer = 99.9
	//fmt.Println(*myFloatPointer)
	//fmt.Println(&myFloatPointer)
	//fmt.Println(myFloat)

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	p := !false
	printPointer(&p)

	var number float64 = 19
	fmt.Println(number)
	double(&number)
	fmt.Println(number)


	truth := true
	negate(&truth)
	fmt.Println(truth)
	lies := false
	negate(&lies)
	fmt.Println(lies)


}


func createPointer() *float64 {
	myFloat := 64.0
	return &myFloat
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func double(numberPointer *float64){
	*numberPointer *= 2
	//fmt.Println(number)
}

func negate(myBoolPointer *bool) bool {
	*myBoolPointer = !*myBoolPointer
	return *myBoolPointer
}