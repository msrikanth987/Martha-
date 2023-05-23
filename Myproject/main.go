package main

import (
	"Myproject/Myproject/helper"
	"fmt"
	"time"
)

var conferenceName string = "Go conference"

const conferenceTickets int = 50

var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstname       string
	lastName        string
	email           string
	numberofTickets uint
}

func main() {

	greetUsers()

	for {
		firstname, lastName, email, userTickets := getuserInput()

		isvaildname, isvailemail, isvaildticketNumbe := helper.Validateuserinput(firstname, lastName, email, userTickets, remainingTickets)

		if isvaildname && isvailemail && isvaildticketNumbe {
			bookTickets(userTickets, firstname, lastName, email)
			sendTicket(userTickets, firstname, lastName, email)
			firstnames := getFirstNames() // printFristname function
			fmt.Printf("these are all our bookings: %v\n", firstnames)
			if remainingTickets == 0 {
				// end program
				fmt.Println("our conference booked out. come back next year")
				break
			}

		} else {
			if !isvaildname {
				fmt.Println("frist name or last name you enter is to small")
			}
			if !isvailemail {
				fmt.Println("email address you enetered doen't contain @ sign")
			}
			if !isvaildticketNumbe {
				fmt.Println("number of tickets you enetred is invaild")
			}

		}

	}

}
func greetUsers() {
	fmt.Printf("welcome to %v booking application\n ", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
func getFirstNames() []string {
	firstnames := []string{}
	for _, booking := range bookings {
		firstnames = append(firstnames, booking.firstname)
	}
	return firstnames
}

func getuserInput() (string, string, string, uint) {
	var firstname string
	var lastName string
	var email string
	var userTickets uint
	// ask user for their Name
	fmt.Println("enter your first name :")
	fmt.Scan(&firstname)

	fmt.Println("enter your last name :")
	fmt.Scan(&lastName)

	fmt.Println("enter your email address:")
	fmt.Scan(&email)

	fmt.Println("enter number of tickets :")
	fmt.Scan(&userTickets)
	return firstname, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstname string, lastName string, email string) {
	remainingTickets = remainingTickets - (userTickets)

	// create a map for a user

	var userData = UserData{
		firstname:       firstname,
		lastName:        lastName,
		email:           email,
		numberofTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("list os bookigs is%v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will recive a conformatiom email at %v\n", firstname, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
func sendTicket(userTickets uint, firstname string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstname, lastName)
	fmt.Println(("*************************************"))
	fmt.Printf("sending ticket:\n %v \n to email addres %v\n", ticket, email)
	fmt.Printf("######################")
}
