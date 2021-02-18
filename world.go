package main

type World struct {
	Hives []Hive
}

func NewWorld() *World {
	w := &World{}
	return w
}
