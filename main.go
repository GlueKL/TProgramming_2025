package main

import "fmt"

type Plane struct {
	Model string
	Seats int
	Speed int
}

func (p Plane) getModel() string {
	return p.Model
}

func (p Plane) getSeats() int {
	return p.Seats
}

func (p Plane) getSpeed() int {
	return p.Speed
}

func (p *Plane) setSpeed(speed int) {
	p.Speed = speed
}

func (p Plane) getPlaneInfo() string {
	return fmt.Sprintf("Model: %s, Seats: %d, Speed: %d", p.getModel(), p.getSeats(), p.getSpeed())
}

func NewPlane(model string, seats int, speed int) Plane {
	return Plane{
		Model: model,
		Seats: seats,
		Speed: speed,
	}
}
func main() {
	fmt.Println("Вариант 9")

	plane1 := NewPlane("Boeing 747", 400, 900)
	plane2 := NewPlane("Airbus A380", 525, 980)
	plane3 := NewPlane("Cessna 172", 4, 100)

	fmt.Println(plane1.getPlaneInfo())
	fmt.Println(plane2.getPlaneInfo())
	fmt.Println(plane3.getPlaneInfo())

	plane1.setSpeed(1000)
	plane2.setSpeed(1100)
	plane3.setSpeed(120)

	fmt.Println(plane1.getPlaneInfo())
	fmt.Println(plane2.getPlaneInfo())
	fmt.Println(plane3.getPlaneInfo())
}
