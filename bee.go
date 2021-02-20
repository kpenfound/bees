package main

import (
	"encoding/json"
	"os"
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

func NewBee(l Location) *Bee {
	b := &Bee{
		Id:           NewId(),
		Nectar:       0,
		TripDuration: 0,
		Life:         BeeLife,
		location:     l,
	}
	return b
}

func (b *Bee) GetJobspec() NomadJob {
	var job NomadJob
	spec := strings.Replace(DefaultJob, "bzzz", b.Id, -1)
	json.Unmarshal([]byte(spec), &job)
	return job
}

func (b *Bee) Spawn(n *NomadAPI, r *Redis) {
	n.CreateJob(b)
	r.SaveBee(*b)
}

func (b *Bee) Die(r *Redis) {
	r.DeleteBee(b)
	//Goodbye
	os.Exit(0)
}

func (b *Bee) Lifecycle() {
	r := NewRedis()
	for {
		see, _ := r.See(b.location, BeeSight, b.Id)
		loc, landed := b.decide(see)
		b.location = loc
		switch landed {
		case FlowerCode:
			f, _ := r.GetFlowerAt(loc)
			f.Pollinate(r)
			b.Nectar++
			break
		case HiveCode:
			h, _ := r.GetHiveAt(loc)
			h.Visit(b.Nectar, r)
			b.Nectar = 0
			b.TripDuration = 0
			break
		}

		b.TripDuration++
		b.Life--

		if b.Life <= 0 || b.TripDuration > BeeTravelLimit {
			b.Die(r)
		} else {
			r.SaveBee(*b)
		}

		time.Sleep(Tick * time.Millisecond)
	}
}

// Heres the fun part
func (b *Bee) decide(vision [][]byte) (Location, byte) {
	target := Location{X: BeeSight * 2, Y: BeeSight * 2}
	relativeLoc := Location{X: BeeSight, Y: BeeSight} // Hmm
	for i := 0; i < len(vision); i++ {
		for j := 0; j < len(vision[i]); j++ {
			switch vision[i][j] {
			case BeeCode:
				break
			case FlowerCode:
				if b.Nectar < BeeNectarCapacity {
					loc := Location{X: i, Y: j}
					if relativeLoc.distance(loc) < relativeLoc.distance(target) {
						target = loc
					}
				}
				break
			case HiveCode:
				if b.Nectar == BeeNectarCapacity {
					loc := Location{X: i, Y: j}
					if relativeLoc.distance(loc) < relativeLoc.distance(target) {
						target = loc
					}
				}
				break
			}
		}
	}
	bearing := relativeLoc.bearing(target)
	b.lastMove = bearing
	move := relativeLoc.moveTo(target)
	landing := vision[move.X][move.Y]
	newX := b.location.X + (move.X - BeeSight)
	newY := b.location.Y + (move.Y - BeeSight)
	return Location{X: newX, Y: newY}, landing
}
