package main

import (
	"Golang-booking-app/validation"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var conferenceName string = "Go Conference"

const totalTickets int = 100

var remainingTickets uint16 = 100

var bookings = make([]map[string]string, 0)

var wg = sync.WaitGroup{}

func main() {

	for remainingTickets > 0 {

		welcomeFunction()

		isInputValid, userData := validation.EnterAndValidateInputs(remainingTickets)

		if !isInputValid {
			fmt.Println("Please enter correct details !!")
			continue
		}

		bookTickets(userData)

		wg.Add(1)
		go sendTicket(userData)

		addToBookingList(userData)

	}

	wg.Wait()

	fmt.Println("All tickets are booking! Please try next time. \n")

}

func welcomeFunction() {

	fmt.Println("Welcome to the", conferenceName, "Booking Centre !!!")
	fmt.Println(`Reserve your seat here !!!`)
	fmt.Printf("We have total seats %v ,only %v seats are left\n", totalTickets, remainingTickets)

}

func bookTickets(data validation.UserData) {

	fmt.Printf("Your %v tickets are booked %v ! Confirmation will be sent shortly.\n",
		data.NumberOftickets, data.FirstName+" "+data.LastName)

	remainingTickets = remainingTickets - data.NumberOftickets

}

func addToBookingList(data validation.UserData) {

	var userData = make(map[string]string)
	userData["firstName"] = data.FirstName
	userData["lastName"] = data.LastName
	userData["email"] = data.EmailId
	userData["tickets"] = strconv.FormatUint(uint64(data.NumberOftickets), 10)

	bookings = append(bookings, userData)

	fmt.Println("People who booked the ticket so far : ")

	for _, booking := range bookings {

		fmt.Printf("%v %v \n", booking["firstName"], booking["lastName"])
	}
}

func sendTicket(data validation.UserData) {

	time.Sleep(50 * time.Second)

	ticket := fmt.Sprintf(" This is email confirmation for your %v %v %v. You have booked %v tickets with us.\n",
		conferenceName, data.FirstName, data.LastName, data.NumberOftickets)

	fmt.Println("###### Sending Ticket ######")
	fmt.Println(ticket)
	fmt.Println("###### Ticket Sent ########")

	wg.Done()
}
