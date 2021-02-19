package main

type Hive struct {
	Id       string
	Nectar   int
	Bees     []*Bee
	location Location
}

func NewHive(id string, l Location) *Hive {
	h := &Hive{
		Id:       id,
		Nectar:   0,
		location: l,
	}
	return h
}

func (h *Hive) Step() {
}

func (h *Hive) Visit(nectar int, r *Redis) {
	h.Nectar += nectar
}

func (h *Hive) SpawnBee(n *NomadAPI, r *Redis) {
	id := NewId()
	b := NewBee(id)
	b.location = h.location
	b.Spawn(n, r)
}

func (h *Hive) SpawnBees(n *NomadAPI, r *Redis, count int) {
	for x := 0; x < count; x++ {
		h.SpawnBee(n, r)
	}
}
