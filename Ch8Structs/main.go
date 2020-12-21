package main

import (
	"fmt"
	"hfgo/Ch8Structs/magazine"
)


//func printInfo(s *student){
//	fmt.Println("Name: ", s.name)
//	fmt.Printf("Grade: %0.1f\n", s.grade)
//}

func main(){

	var address magazine.Address
	address.Street = "123 Oak St."
	address.City = "Omaha"
	address.State ="NE"
	address.PostalCode = "68111"
	fmt.Println(address)

	sub := magazine.Subscriber{Name: "Jonny"}
	sub.HomeAddress = address
	//sub.HomeAddress.Street = "123 Oak St"
	//sub.HomeAddress.City = "Omaha"
	//sub.HomeAddress.State = "NE"
	//sub.HomeAddress.PostalCode = "68111"
	sub.Rate = 4.99
	fmt.Println(sub)

	var emp magazine.Employee
	emp.Name = "Joe"
	emp.Salary = 60000
	emp.Street = "456 Elm St"
	emp.Address.City = "Portland"
	emp.State = "OR"
	emp.Address.PostalCode = "97222"
	fmt.Println(emp)

}