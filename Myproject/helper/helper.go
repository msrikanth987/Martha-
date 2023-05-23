package helper

import "strings"

func Validateuserinput(firstname string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isvaildname := len(firstname) >= 2 && len(lastName) >= 2
	var isvailemail = strings.Contains(email, "@")
	var isvaildticketNumbe = userTickets > 0 && userTickets <= remainingTickets
	return isvaildname, isvailemail, isvaildticketNumbe
}
