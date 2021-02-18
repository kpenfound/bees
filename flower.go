package main

type Flower struct {
	Id               string
	Pollinated       bool
	Life             uint64
	PollinateCounter uint64
	location         Location
}

func NewFlower(id string) *Flower {
	f := &Flower{
		Id:               id,
		Pollinated:       false,
		Life:             FlowerLife,
		PollinateCounter: 0,
	}
	return f
}

func (f *Flower) Step() {
	f.Life--
}

func (f *Flower) Pollinate() {
	f.Pollinated = true
	f.PollinateCounter++
}
