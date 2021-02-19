package main

import (
	"encoding/json"
	"strings"
	"time"
)

type Bee struct {
	Id           string
	Nectar       int
	TripDuration int
	Life         int
	location     Location
	lastMove     int
}

func NewBee(id string) *Bee {
	b := &Bee{
		Id:           id,
		Nectar:       0,
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

func (b *Bee) step() {
	b.TripDuration++
	b.Life--
}

func (b *Bee) Spawn(n *NomadAPI, r *Redis) {
	n.CreateJob(b)
	r.SaveBee(*b)
}

func (b *Bee) Die() {
	n := NewNomad()
	n.DeleteJob(b)
}

func (b *Bee) Lifecycle() {
	r := NewRedis()
	for {
		see := r.See(b.location, BeeSight)
		loc, landed := b.decide(see)
		b.location = loc
		switch landed {
		case FlowerCode:
			f := r.GetFlowerAt(loc)
			f.Pollinate(r)
			r.SaveFlower(f)
			b.Nectar++
			break
		case HiveCode:
			h := r.GetHiveAt(loc)
			h.Visit(b.Nectar, r)
			r.SaveHive(h)
			b.Nectar = 0
			b.TripDuration = 0
			break
		}

		r.SaveBee(*b)
		time.Sleep(Tick * time.Millisecond)
	}
}

// Heres the fun part
func (b *Bee) decide(vision [][]byte) (Location, byte) {
	for i := 0; i < len(vision); i++ {
		for j := 0; j < len(vision[i]); j++ {
			switch vision[i][j] {
			case BeeCode:
				break
			case FlowerCode:
				break
			case HiveCode:
				break
			}
		}
	}
	b.lastMove = 0
	return Location{X: 0, Y: 0}, EmptyCode
}
