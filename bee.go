package main

type Bee struct {
	Id string
}

func NewBee(id string) *Bee {
	b := &Bee{Id: id}
	return b
}
