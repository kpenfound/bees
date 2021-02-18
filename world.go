package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/go-redis/redis/v8"
)

type World struct {
	Hives []*Hive
	Flowers []*Flower
	redis *redis.Client
	nomad *NomadAPI
}

func NewWorld() *World {
	w := &World{}

	logger := hclog.New(nil)
	w.nomad = NewNomad(logger)
	w.redis = NewRedisClient()

	for x := 0; x <= WorldX; x++ {
		for y := 0; y <= WorldY; y++ {
			if x % (WorldX / WorldStartingFlowers) == 0 && y % (WorldY / WorldStartingFlowers) == 0 {
				f := NewFlower(NewId(), Location{X:x,Y:y})
				w.Flowers = append(w.Flowers, f)
			}

			if x % (WorldX / WorldStartingHives) == 0 && y % (WorldY / WorldStartingHives) == 0 {
				h := NewHive(NewId(), Location{X:x,Y:y})
				w.Hives = append(w.Hives, h)
			}
		}
	}
	return w
}

func (w *World) Simulate() {
	// Do simulation
}

func (w *World) Step() {
}

type Location struct {
	X int
	Y int
}

func NewId() string {
	return "a"
}
