package main

type World struct {
	Hives []*Hive
}

func NewWorld() *World {
	w := &World{}
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
