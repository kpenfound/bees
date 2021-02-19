package main

type Flower struct {
	Id               string
	Pollinated       bool
	Life             int
	PollinateCounter int
	location         Location
}

func NewFlower(id string, l Location) *Flower {
	f := &Flower{
		Id:               id,
		Pollinated:       false,
		Life:             FlowerLife,
		PollinateCounter: 0,
		location:         l,
	}
	return f
}

func (f *Flower) step() {
	f.Life--
}

func (f *Flower) Spawn() {
	// Save
}

func (f *Flower) Pollinate(r *Redis) {
	f.Pollinated = true
	f.PollinateCounter++
}

func (f *Flower) Die() {
	// Delete
}
