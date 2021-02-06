package main

import (
	"fmt"
)

type contact struct {
	firstName string
	lastName  string
	number    int64
}

func newcontact(firstName, lastName string, number int64) contact {
	return contact{
		firstName: firstName,
		lastName:  lastName,
		number:    number,
	}
}

//methods
func (c *contact) updateNumber(number int64) {
	c.number = number
}

func (c *contact) fullName() string {
	return fmt.Sprintf("%s %s", c.firstName, c.lastName)
}

func (c *contact) getContact() string {
	return fmt.Sprintf("Fist Name: %s\nLast Name: %s\nPhone Number: %d\n", c.firstName, c.lastName, c.number)
}

var contactList = []contact{
	newcontact("Parpiyev", "Mirobidjon", 944344143),
	newcontact("Habibullayev", "Begali", 905286368),
	newcontact("Abdumalikov", "Halilullo", 999014533),
	newcontact("Yuldashev", "Asadbek", 938554784),
}

// func main() {
// 	for _, i := range contactList {
// 		fmt.Println(i.getContact())
// 	}
// }
