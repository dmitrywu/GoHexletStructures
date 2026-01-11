package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
}

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount float64) bool {
	if a.Balance >= amount {
		a.Balance -= amount
		return true
	}
	return false
}

func (u User) Greet() string {
	return "Hello, " + u.Name
}

type Draft struct {
	Title     string
	Items     []string
	published bool
}

func (d *Draft) AddItem(item string) {
	// Он добавляет непустой item в Items, пустые строки игнорируются.
	if item != "" {
		d.Items = append(d.Items, item)
	}
}

func (d *Draft) Publish() {
	// переводит черновик в опубликованный статус (published = true).
	d.published = true
}

func (d *Draft) IsPublished() bool {
	// возвращающий текущее значение флага. Получатель по значению, чтобы метод корректно работал у копий структуры.
	return d.published
}

func main() {
	fmt.Print("\033[H\033[2J")

	draft := Draft{Title: "Winter combo"}
	draft.AddItem("Tea")
	draft.AddItem("")
	draft.AddItem("Mug")
	draft.Publish()

	fmt.Println(draft.Items) // []string{"Tea", "Mug"}
	draft.IsPublished()      // true

	copyDraft := draft
	copyDraft.IsPublished() // true

	acc := Account{Owner: "Maria", Balance: 100}
	acc.Deposit(50)
	fmt.Println(acc.Balance)

	ok := acc.Withdraw(200)
	if ok {
		fmt.Println("Успешно")
	} else {
		fmt.Println("Ошибочка")
	}

	user := User{Name: "Иван", Email: "ivan@example.com"}
	fmt.Println(user.Greet()) // "Привет, Иван"

}
