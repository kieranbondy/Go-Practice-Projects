package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Henderson",
		contactInfo: contactInfo{
			zip:   10003,
			email: "jh@gmail.com",
		},
	}
	jim.updateName("Tom")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}
func (p *person) updateName(newName string) {
	(*p).firstName = newName
}
