package main

import "fmt"

type Logger struct{}

func (l Logger) Log(message string) {
	fmt.Println("LOG:", message)
}

type Service struct {
	Name   string
	Logger // встраиваем логгер

}

func main() {
	fmt.Print("\033[H\033[2J")
	s := Service{Name: "OrderService"}
	s.Log("Запущен сервис") // метод Logger стал методом Service
}
