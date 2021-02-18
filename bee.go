package main

type Bee struct {
	Id           string
	HasNectar    bool
	TripDuration uint64
	Life         uint64
	location     Location
}

func NewBee(id string) *Bee {
	b := &Bee{
		Id:           id,
		HasNectar:    false,
		TripDuration: 0,
		Life:         BeeLife,
	}
	return b
}

func (b *Bee) Step() {
	b.TripDuration++
	b.Life--
}
