package main

type Flower struct {
	Id               string
	PollinateCounter int
	location         Location
}

func NewFlower(l Location) *Flower {
	f := &Flower{
		Id:               NewId(),
		PollinateCounter: 0,
		location:         l,
	}
	return f
}

func (f *Flower) Spread(r *Redis) {
	newLoc := findFreeNeighbor(f.location, r)
	newf := NewFlower(newLoc)
	r.SaveFlower(*newf, true)
}

func findFreeNeighbor(loc Location, r *Redis) Location {
	// TODO
	neighbor := Location{X: -1, Y: -1}
	return neighbor
}

func (f *Flower) Pollinate(r *Redis) {
	f.Spread(r)
	f.PollinateCounter++
	if f.PollinateCounter >= FlowerPollinateLimit {
		f.Die(r)
	} else {
		r.SaveFlower(*f, false)
	}
}

func (f *Flower) Die(r *Redis) {
	r.DeleteFlower(f)
}
