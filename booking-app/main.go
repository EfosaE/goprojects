package main
import "fmt"


func main () {
	const conferenceName = "Track Fest"
	const ticketsAvailable = 50

	fmt.Printf("Welcome to %v 😁, We have %v tickets available for you 😊. \n", conferenceName, ticketsAvailable)

	var userName string
	var email string
	var ticketsBought uint

	fmt.Print("Enter your name:")
	fmt.Scan(&userName)

	fmt.Print("Enter your email:")
	fmt.Scan(&email)

	fmt.Print("Enter the desired number of tickets:")
	fmt.Scan(&ticketsBought)

	fmt.Printf("%v @ %v bought %v",userName, email, ticketsBought)
}