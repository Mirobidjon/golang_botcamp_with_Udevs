package main

import (
	"fmt"
)

type contact struct {
	firstName string
	lastName  string
	number    int
}

func newContact(firstName, lastName string, number int) contact {
	return contact{
		firstName: firstName,
		lastName:  lastName,
		number:    number,
	}
}

//methods
func (c *contact) updateNumber(number int) {
	c.number = number
}

func (c *contact) fullName() string {
	return fmt.Sprintf("%s %s", c.firstName, c.lastName)
}

func (c *contact) getContact() string {
	return fmt.Sprintf("Fist Name: %s\nLast Name: %s\nPhone Number: %d\n", c.firstName, c.lastName, c.number)
}

var contactList = []contact{
	newContact("Parpiyev", "Mirobidjon", 944344143),
	newContact("Habibullayev", "Begali", 905286368),
	newContact("Abdumalikov", "Halilullo", 999014533),
	newContact("Yuldashev", "Asadbek", 938554784),
}

func main() {
	for _, i := range contactList {
		fmt.Println(i.getContact())
	}
}
