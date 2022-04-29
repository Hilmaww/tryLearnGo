package main

import "fmt"

func main(){
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Println("Welcome %v to our conference booking app", conferenceName)
	fmt.Println("We have %v and %v tickets left", conferenceTickets, remainingTickets )
	fmt.Println("Get your tickets here to attend")


}