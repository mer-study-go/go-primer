package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	mer := person{
		firstName: "Mer",
		lastName:  "JQ",
		contact: contactInfo{
			email:   "mer@gmail.com",
			zipCode: 94000,
		},
	}
	mer.updateName("Jieqiong")
	mer.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}
