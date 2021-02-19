package main

import (
	"encoding/json"
	"strings"
)

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

func (b *Bee) GetJobspec() NomadJob {
	var job NomadJob
	spec := strings.Replace(DefaultJob, "bzzz", b.Id, -1)
	json.Unmarshal([]byte(spec), &job)
	return job
}

func (b *Bee) Step() {
	b.TripDuration++
	b.Life--
}

func (b *Bee) Spawn() {
	n := NewNomad()
	n.CreateJob(b)
}

func (b *Bee) Die() {
	n := NewNomad()
	n.DeleteJob(b)
}
