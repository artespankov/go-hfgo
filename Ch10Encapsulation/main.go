package main

import (
	"hfgo/Ch10Encapsulation/calendar"
	"fmt"
	"log"
)




func main(){
	event := calendar.Event{}
	err := event.SetTitle("NY Eve")
	if err != nil{
		log.Fatal(err)
	}
	err = event.SetYear(2020)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetMonth(12)
	if err != nil {
		log.Fatal(err)
	}
	err = event.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %d %d %d\n", event.Title(), event.Year(), event.Month(), event.Day())
}
