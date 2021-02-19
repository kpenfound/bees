package main

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
		h.SpawnBee(n)
	}
	r.SaveHive(*h)
}

func (h *Hive) SpawnBee(n *NomadAPI) {
	b := NewBee(h.location)
	b.Spawn(n)
}

func (h *Hive) SpawnBees(n *NomadAPI, count int) {
	for x := 0; x < count; x++ {
		h.SpawnBee(n)
	}
}
