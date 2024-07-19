package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)
var conferenceName = "Go Conference"
const conferenceTickets uint = 50
var remainingTickets uint = 50
var bookings = make([]userData,0)

type userData struct {
	firstName string
	lastName string
	email string
	numTickets uint
}
var wg = sync.WaitGroup{}

func main() {

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	
	greetUsers()
	validateInput(firstName,lastName,email,userTickets)
	wg.Wait()
	}
func greetUsers()  {
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("We have a total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend")
	}

func sendTickets(userTickets uint, firstName string, lastName string, email string)  {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	fmt.Printf("Sending Ticket:\n %v\n to email address %v\n",  ticket, email)
	fmt.Println("-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-")
	wg.Done()
}
func validateInput(fName string, lName string, email string, uTickets uint) {

	fmt.Println("Enter your first name:")
	fmt.Scanln(&fName)


	fmt.Println("Enter your last name:")
	fmt.Scanln(&lName)

	fmt.Println("Enter your email:")
	fmt.Scanln(&email)

	fmt.Println("Enter the number of tickets you want:")
	fmt.Scanln(&uTickets)

	isValidName := len(fName)>=2 && len(lName)>=2
	isValidEmail := strings.Contains(email,"@")
	isValidTicketNumber := uTickets>0 && uTickets <= remainingTickets

	if isValidName && isValidEmail && isValidTicketNumber{
		remainingTickets = remainingTickets - uTickets

		var userData = userData{
			firstName: fName,
			lastName: lName,
			email: email,
			numTickets: uTickets,
		}

		bookings = append(bookings, userData)
		fmt.Printf("List of bookings is %v\n",bookings)
	
		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", fName, lName, uTickets, email)
		fmt.Printf("%v tickets are remaining for the %v\n", remainingTickets, conferenceName)

		wg.Add(1)
		go sendTickets(uTickets, fName,lName,email)
	
		fNames := []string{}
		for _, booking := range bookings {
				fNames = append(fNames, booking.firstName)
			}
			fmt.Printf("The first names of all our bookings are: %v\n",fNames)

		
		if remainingTickets == 0 {
			fmt.Printf("%v is sold out. Come back next year\n", conferenceName)
		}
	} else {
		if !isValidName {
			fmt.Println("Your firstname or lastname entered is too short")
		}
		if !isValidEmail {
			fmt.Println("You entered an invalid e-mail address")
		}
		if !isValidTicketNumber {
			fmt.Println("You entered an invalid ticket number")
			fmt.Println("")
		}
	}
}