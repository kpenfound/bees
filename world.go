package main

import (
	"math"

	"github.com/google/uuid"
)

func NewWorld() {
	n := NewNomad()
	r := NewRedis()

	for x := 0; x <= WorldX; x++ {
		for y := 0; y <= WorldY; y++ {
			if x%(WorldX/WorldStartingFlowers) == 0 && y%(WorldY/WorldStartingFlowers) == 0 {
				f := NewFlower(Location{X: x, Y: y})
				r.SaveFlower(*f)
			}

			if x%(WorldX/WorldStartingHives) == 0 && y%(WorldY/WorldStartingHives) == 0 {
				h := NewHive(Location{X: x, Y: y})
				h.SpawnBees(n, r, HiveStartingBees)
				r.SaveHive(*h)
			}
		}
	}
}

type Location struct {
	X int
	Y int
}

func (a Location) distance(b Location) int {
	dX := float64(b.X - a.X)
	dY := float64(b.Y - a.Y)
	return int(math.Sqrt(dX*dX + dY*dY))
}

func (a Location) bearing(b Location) int {
	dX := float64(b.X - a.X)
	dY := float64(b.Y - a.Y)
	return int(math.Atan2(dY, dX))
}

func (a Location) moveTo(b Location) Location {
	dX := float64(b.X - a.X)
	dY := float64(b.Y - a.Y)
	if math.Abs(dX) > math.Abs(dY) {
		a.X += int(dX / math.Abs(dX))
	} else {
		a.Y += int(dY / math.Abs(dY))
	}
	return a
}

func NewId() string {
	return uuid.NewString()
}
