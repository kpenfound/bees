package main

import "fmt"

type Hive struct {
	Id       string
	Nectar   int
	location Location
}

func NewHive(l Location) *Hive {
	h := &Hive{
		Id:       NewId(),
		Nectar:   0,
		location: l,
	}
	return h
}

func (h *Hive) Visit(nectar int, r *Redis) {
	h.Nectar += nectar
	if h.Nectar >= BeeNectarCost {
		h.Nectar -= BeeNectarCost
		n := NewNomad()
		h.SpawnBee(n, r)
	}
	r.SaveHive(*h, false)
}

func (h *Hive) SpawnBee(n *NomadAPI, r *Redis) {
	b := NewBee(h.location)
	b.Spawn(n, r)
	fmt.Printf("Bee spawned at %d %d\n", h.location.X, h.location.Y)
}

func (h *Hive) SpawnBees(n *NomadAPI, r *Redis, count int) {
	for x := 0; x < count; x++ {
		h.SpawnBee(n, r)
	}
}
