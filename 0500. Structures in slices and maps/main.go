package main

import "fmt"

type Order struct {
	ID       int
	Customer string
	Items    []string
	Status   string
}

type Student struct {
	Name string
	Age  int
}

type Shipment struct {
	ID        string
	City      string
	Items     []string
	Delivered bool
}

func FilterByCity(shipments []Shipment, city string) []Shipment {
	//возвращающую новый срез отправлений в указанном городе (порядок сохраняется).
	cityFiltered := make([]Shipment, 0, len(shipments))
	for _, s := range shipments {
		if s.City == city {
			cityFiltered = append(cityFiltered, s)
		}
	}
	return cityFiltered
}

func DeliveredIDs(shipments []Shipment) []string {
	//которая собирает ID доставленных отправлений в отдельный срез.
	deli := make([]string, 0, len(shipments))
	for _, s := range shipments {
		if s.Delivered {
			deli = append(deli, s.ID)
		}
	}
	return deli
}

func IndexByID(shipments []Shipment) map[string]Shipment {
	// которая строит карту ID -> Shipment. Последние значения по одинаковым ключам перезаписывают предыдущие.
	idx := make(map[string]Shipment)
	for _, s := range shipments {
		idx[s.ID] = s
	}
	return idx
}

func main() {
	fmt.Print("\033[H\033[2J")

	shipments := []Shipment{
		{ID: "A1", City: "Berlin", Delivered: true},
		{ID: "B2", City: "Paris", Delivered: false},
		{ID: "C3", City: "Berlin", Delivered: true},
	}

	fmt.Println(FilterByCity(shipments, "Berlin"))
	fmt.Println(DeliveredIDs(shipments))
	fmt.Println(IndexByID(shipments)["B2"])

	orders := []Order{
		{ID: 101, Customer: "Иван", Items: []string{"Ноутбук", "Мышь"}, Status: "new"},
		{ID: 102, Customer: "Мария", Items: []string{"Телефон"}, Status: "processing"},
	}

	for _, order := range orders {
		fmt.Printf("Заказ %d для %s, статус: %s\n", order.ID, order.Customer, order.Status)
	}

	ordersMap := map[int]Order{
		101: {ID: 101, Customer: "Иван", Items: []string{"Ноутбук", "Мышь"}, Status: "new"},
		102: {ID: 102, Customer: "Мария", Items: []string{"Телефон"}, Status: "processing"},
	}

	order := ordersMap[101]
	fmt.Println(order.Customer, order.Items) // Иван

	grades := map[Student]float64{
		{Name: "Ольга", Age: 20}: 4.5,
		{Name: "Петр", Age: 22}:  3.9,
	}
	fmt.Println(grades)

}
