package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{} // slice

	fmt.Printf(
		"conferenceTickets is %v, remainingTicket is %v, conferenceName is %v\n",
		conferenceTickets,
		remainingTickets,
		conferenceName,
	)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf(
		"We have total of %v tickets and  %v are still available\n",
		conferenceTickets,
		remainingTickets,
	)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		fmt.Println("your first name:")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address:")
		fmt.Scan(&email)

		fmt.Println("Enter your number of tickets:")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			// booking logic
			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf(
				"Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n",
				firstName,
				lastName,
				userTickets,
				email,
			)

			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
			fmt.Printf("These are all our bookings: %v\n", bookings)

			firstNames := []string{} // create a blank slice

			for _, booking := range bookings {
				var names = strings.Fields(booking)                             // split the string into a slice of strings
				firstNames = append(firstNames, names[0])                       // append the first name to the firstNames slice }
				fmt.Printf("The first names of bookings are: %v\n", firstNames) // print the first names
			}

			if remainingTickets == 0 {
				// if the remaining tickets is 0 then we break out of the loop
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

		} else {

			if !isValidName {
				fmt.Println("first name or last name you entered is too short")
			}

			if !isValidEmail {
				fmt.Println("email address you entered doesn't contain @ sign")
			}

			if !isValidTicketNumber {
				fmt.Println("number of tickets you entered is invalid")
			}

			fmt.Println("Your input data is invalid. Please try again.")

		}

	}
}
