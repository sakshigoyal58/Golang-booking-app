package validation

import (
	"fmt"
	"strings"
)

type UserData struct {
	FirstName       string
	LastName        string
	EmailId         string
	PhoneNumber     uint64
	NumberOftickets uint16
}

func EnterAndValidateInputs(remainingTickets uint16) (bool, UserData) {

	var newUserData UserData

	fmt.Println("Please enter your FirstName and LastName :")
	fmt.Scan(&newUserData.FirstName, &newUserData.LastName)

	if len(newUserData.FirstName) <= 2 || len(newUserData.LastName) <= 2 {
		return false, newUserData
	}

	fmt.Println("Please enter your Email Id :")
	fmt.Scan(&newUserData.EmailId)

	if !strings.Contains(newUserData.EmailId, "@") || !strings.Contains(newUserData.EmailId, ".") {
		return false, newUserData
	}

	fmt.Println("Please enter your Phone number where we can send confirmation :")
	fmt.Scan(&newUserData.PhoneNumber)

	if newUserData.PhoneNumber <= 999 {
		return false, newUserData
	}

	fmt.Println("Numbers of tickets Needs to be booked : ")
	fmt.Scan(&newUserData.NumberOftickets)

	if newUserData.NumberOftickets <= 0 || newUserData.NumberOftickets > remainingTickets {
		fmt.Printf("Sorry! we cannot book these many tickets! only %v ticketa are left.\n", remainingTickets)
		return false, newUserData
	}

	return true, newUserData

}
