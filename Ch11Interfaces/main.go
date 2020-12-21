package main

import (
	"fmt"
	"hfgo/Ch11Interfaces/gadget"
)

func playList(device gadget.Player, songs []string){
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}
func tryOut(player gadget.Player){
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

type Robot string
func (r Robot) MakeSound(){
	fmt.Println("Beep Boop")
}
func (r Robot) Walk(){
	fmt.Println("Powering legs")
}

type NoiseMaker interface {
	MakeSound()
}

func main() {
	var player gadget.Player = gadget.TapePlayer{}
	mixtape := []string{"Jessie's Girl", "Whip It", "9 to 5"}
	playList(player, mixtape)

	player = gadget.TapeRecorder{}
	mixtape = []string{"Chika Bonita", "Highway to Hell", "Enter Sandman"}
	playList(player, mixtape)


	tryOut(gadget.TapePlayer{})
	tryOut(gadget.TapeRecorder{})

	var noiseMaker NoiseMaker = Robot("Botco Ambler")
	noiseMaker.MakeSound()
	var robot Robot = noiseMaker.(Robot)
	robot.Walk()
}