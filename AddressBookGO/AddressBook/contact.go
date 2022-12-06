package addressbook

import ()

type Contact struct {
	FirstName string
	LastName  string
	Email     string
	Phone     string
}

func New(firstname, lastname, email, phone string) Contact {
	c := Contact{
		firstname, lastname, email, phone,
	}
	return c
}
