package main

import "fmt"

type Parcel struct {
	ID          string
	Destination string
	Weight      float64
	Insured     bool
}

func Describe(p Parcel) string {
	checkInsured := "yes"
	if p.Insured == false {
		checkInsured = "no"
	}
	return fmt.Sprintf("#%s -> %s (%.1f kg, insured: %s)", p.ID, p.Destination, p.Weight, checkInsured)
}

func main() {
	fmt.Print("\033[H\033[2J")

	parcel1 := Parcel{ID: "BR-7", Destination: "Berlin", Weight: 12, Insured: true}
	// "#BR-7 -> Berlin (12.0 kg, insured: yes)"
	//t1 := fmt.Sprintf("%s -> %s (%f.1, %t )", parcel1.ID, parcel1.Destination, parcel1.Weight, parcel1.Insured)
	fmt.Println(Describe(parcel1))
	//fmt.Println(t1)

	parcel2 := Parcel{ID: "LK-54", Destination: "Lima", Weight: 0.5, Insured: false}
	//Describe(parcel2) // "#LK-54 -> Lima (0.5 kg, insured: no)"
	fmt.Println(Describe(parcel2))

}
