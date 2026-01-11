package main

import "fmt"

type Health struct {
	Status    string
	Incidents []string
}

type Service struct {
	Name string
	Health
}

func (h *Health) AddIncident(incident string) {
	if incident == "" {
		return
	}

	h.Incidents = append(h.Incidents, incident)

	if h.Status == "" {
		h.Status = "degraded"
	}
}

func (h Health) IsStable() bool {
	if h.Status == "ok" && len(h.Incidents) < 1 {
		return true
	}
	//возвращает true, когда статус "ok" и нет инцидентов.
	return false
}

func NewService(name string) Service {
	// создаёт сервис со статусом ok и пустым списком инцидентов.
	return Service{
		Name: name,
		Health: Health{
			Status:    "ok",
			Incidents: nil,
		},
	}
}

func main() {
	fmt.Print("\033[H\033[2J")

}
