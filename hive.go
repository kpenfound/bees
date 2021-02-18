package main

type Hive struct {
	Id       string
	Nectar   uint64
	Bees     []*Bee
	location Location
}

func NewHive(id string) *Hive {
	h := &Hive{Id: id}
	return h
}

func (h *Hive) Step() {
}
