package main

import "fmt"

type Customer struct {
	Name  string
	Phone string
}

type Address struct {
	City   string
	Street string
	House  string
}

type Delivery struct {
	ID       string
	Customer Customer
	Address  Address
	Items    []string
}

func FormatDelivery(d Delivery) string {
	// #<ID> <Name> (<Phone>) -> <City>, <Street> <House>.
	// "#10 Ivan (+7999) -> Moscow, Lenina 10"
	t := fmt.Sprintf("#%s %s (%s) -> %s, %s %s.", d.ID, d.Customer.Name, d.Customer.Phone, d.Address.City, d.Address.Street, d.Address.House)
	return t
}

func main() {
	fmt.Print("\033[H\033[2J")

	d := Delivery{
		ID:       "10",
		Customer: Customer{Name: "Ivan", Phone: "+7999"},
		Address:  Address{City: "Moscow", Street: "Lenina", House: "10"},
		Items:    []string{"Tea", "Cookies"},
	}

	// FormatDelivery(d) // "#10 Ivan (+7999) -> Moscow, Lenina 10"
	fmt.Println(FormatDelivery(d))

}
