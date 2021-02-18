package main

type Hive struct {
	Id       string
	Nectar   uint64
	Bees     []*Bee
	location Location
}

func NewHive(id string, l Location) *Hive {
	h := &Hive{
		Id: id,
		Nectar: 0,
		location: l,
	}
	return h
}

func (h *Hive) Step() {
}

func (h *Hive) SpawnBee(n NomadAPI) {
	id := NewId()
	b := NewBee(id)
	n.CreateJob(b)
}

func (h *Hive) SpawnBees(n NomadAPI, count int) {
	for x := 0; x < count; x++ {
		h.SpawnBee(n)
	}
}
