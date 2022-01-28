package main

import (
	"fmt"
	"rivercrossing/event"
)

func main() {
	fmt.Println(event.ViewEvent())

	event.BåtVestKylling()
	fmt.Println(event.ViewEvent())

	event.BåtØstKylling()
	fmt.Println(event.ViewEvent())

	event.LandØstKylling()
	fmt.Println(event.ViewEvent())

	event.BåtTilVest()
	fmt.Println(event.ViewEvent())
}
