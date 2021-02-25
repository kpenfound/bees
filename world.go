package main

import (
	"fmt"
	"math"
	"math/rand"
	"strings"

	"github.com/google/uuid"
)

func NewWorld() {
	n := NewNomad()
	r := NewRedis()

	for i := 0; i < WorldStartingFlowers; i++ {
		x := rand.Intn(WorldX)
		y := rand.Intn(WorldY)
		f := NewFlower(Location{X: x, Y: y})
		r.SaveFlower(*f, true)
		fmt.Printf("Created flower at %d %d\n", x, y)
	}

	for j := 0; j < WorldStartingHives; j++ {
		x := rand.Intn(WorldX)
		y := rand.Intn(WorldY)
		h := NewHive(Location{X: x, Y: y})
		h.SpawnBees(n, r, HiveStartingBees)
		r.SaveHive(*h, true)
		fmt.Printf("Created hive at %d %d\n", x, y)
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
	id := uuid.NewString()
	return strings.Replace(id, "-", "", -1)
}
