package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	var alex = person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "aa@gmail.com",
			zipCode: 82001,
		},
	}

	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	alex.print()
	alex.updateName("Al")
	alex.print()

	s := []string{"One", "Two", "Three"}
	fmt.Println(s)
	updateSlice(s)
	fmt.Println(s)
	//fmt.Printf("%p\n", s)
	//i := int(42)
	//fmt.Printf("1. main  -- i: &i=%p i=%v\n", &i, i)
	name := "Bill"
	namePtr := &name

	fmt.Println(&namePtr)
	printPointer(namePtr)
}

func printPointer(namePointer *string) {
	fmt.Println(&namePointer)
}

func (p *person) updateName(name string) {
	p.firstName = name
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func updateSlice(s []string) {
	s[0] = "Gog"
}
