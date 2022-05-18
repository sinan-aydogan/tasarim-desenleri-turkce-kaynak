package main

type employee struct {
	firstName string
	lastName  string
}

func NewEmployee(firstName string, lastName string) *employee {
	return &employee{firstName, lastName}
}

func (e *employee) clone() cloneablePrototype {
	return &employee{e.firstName, e.lastName}
}
