package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type InvoiceDTO struct {
	ID       int     `json:"id"`
	Customer string  `json:"customer"`
	Amount   float64 `json:"amount"`
	Comment  string  `json:"comment,omitempty"`
	Secret   string  `json:"-"`
}

func EncodeInvoice(inv InvoiceDTO) ([]byte, error) {
	// которая оборачивает json.Marshal
	return json.Marshal(inv)
}

func DecodeInvoice(payload []byte) (InvoiceDTO, error) {
	// оборачивает json.Unmarshal. При ошибке верните нулевую структуру и ошибку.
	var s InvoiceDTO
	e := json.Unmarshal(payload, &s)
	if e != nil {
		return InvoiceDTO{}, fmt.Errorf("Errorочка")
	}
	return s, nil
}

func main() {
	fmt.Print("\033[H\033[2J")

	u := User{Name: "Иван", Age: 30}
	b, v := json.Marshal(u)
	fmt.Println(string(b), v) // {"name":"Иван","age":30}

}
