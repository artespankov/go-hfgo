package main

import "fmt"

type Liters float64
func (l Liters) toGallons() Gallons{
	return Gallons(l * 0.264)
}
type Milliliters float64
func (m Milliliters) toGallons() Gallons{
	return Gallons(m * 0.000264)
}
type Gallons float64


type Number int
func (n *Number) Display(){
	fmt.Println(*n)
}
func (n *Number ) Double(){
	*n *= 2
}


type Title string
func (t Title) sayHi(){
	fmt.Printf("Hi %s!\n", t)
}


func main(){

	name := Title("Art")
	name.sayHi()

	var carFuel Gallons
	var busFuel Liters
	carFuel = Gallons(120.15)
	busFuel = Liters(99.17)
	fmt.Println(carFuel, busFuel)

	number := Number(5)
	number.Double()
	number.Display()

	soda := Liters(2)
	fmt.Printf("%0.2f liters equals %0.2f gallons\n", soda, soda.toGallons())
	water := Milliliters(500)
	fmt.Printf("%0.2f milliliters equals %0.2f gallons", water, water.toGallons())


}



