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

type listtest []string

func main() {

	alex := person{
		firstName: "Alex",
		lastName:  "Chatterjee",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 713144,
		},
	}
	alex.print()
	alexPointer := &alex
	alexPointer.updateName("jimmy")
	alexPointer.print()

	var listRef = listtest{"hello", "hi"}
	updateName(listRef)
	fmt.Println(listRef)
	listRef.updtListref()
	fmt.Println(listRef)

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
func updateName(a []string) {
	a[0] = "test"
}
func (test listtest) updtListref() {
	test[0] = "hello"
}
