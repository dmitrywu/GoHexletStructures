package main

import "fmt"

type Rack struct {
	ID          string
	Temperature float64
	PowerOff    bool
}

func Adjust(r *Rack, delta float64) bool {
	if r == nil {
		return false
	}
	r.Temperature = r.Temperature + delta
	return true
}

func EmergencyShutdown(r *Rack) {
	if r != nil {
		r.PowerOff = true
	}
}

func Snapshot(r *Rack) Rack {
	if r == nil {
		return Rack{}
	}
	return *r
}

func main() {
	fmt.Print("\033[H\033[2J")

}
