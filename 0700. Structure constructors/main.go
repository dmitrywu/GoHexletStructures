package main

import (
	"fmt"
	"time"
)

type Order struct {
	ID     int
	Items  []string
	Status string
}

func NewOrder(id int, items []string) *Order {
	return &Order{
		ID:     id,
		Items:  items,
		Status: "new",
	}
}

type Account struct {
	id      int
	owner   string
	balance float64
}

type Config struct {
	Path    string
	Timeout time.Duration
	Limit   int
}

func NewConfig(path string) *config {
	return &Config{
		Path:    path,
		Timeout: 30 * time.Second,
		Limit:   100,
	}
}

type ReportConfig struct {
	Title           string
	Limit           int
	IncludeArchived bool
}

func NewReportConfig(title string, limit int, includeArchived bool) (ReportConfig, error) {
	if title == "" {
		return ReportConfig{}, fmt.Errorf("Empty")
	}
	if limit <= 0 {
		limit = 100
	}
	return ReportConfig{
		Title:           title,
		Limit:           limit,
		IncludeArchived: includeArchived,
	}, nil
}

func (cfg ReportConfig) MaxItems() int {
	return cfg.Limit
}

func main() {
	fmt.Print("\033[H\033[2J")

}
