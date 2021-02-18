package main

type Flower struct {
	Id string
}

func NewFlower(id string) *Flower {
	f := &Flower{Id: id}
	return f
}
