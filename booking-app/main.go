package main

import "fmt"

func main() {
	const conferenceName = "Track Fest"
	var ticketsAvailable uint = 50
	bookings := []string{}

	conferenceTickets := 50

	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, ticketsAvailable)

	for {
		var userName string
		var email string
		var ticketsBought uint

		// asking for user input
		fmt.Print("Enter your name:")
		fmt.Scan(&userName)

		fmt.Print("Enter your email:")
		fmt.Scan(&email)

		fmt.Print("Enter the desired number of tickets:")
		fmt.Scan(&ticketsBought)

		// book ticket in system
		ticketsAvailable = ticketsAvailable - ticketsBought
		bookings = append(bookings, userName)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, ticketsBought, email)
		fmt.Printf("%v tickets remaining for %v\n", ticketsAvailable, conferenceName)
		fmt.Println(bookings)

		// // print only first names
		// firstNames := []string{}
		// for _, booking := range bookings {
		// 	var names = strings.Fields(booking)
		// 	firstNames = append(firstNames, names[0])
		// }
		// fmt.Printf("The first names %v\n", firstNames)
	}

}
