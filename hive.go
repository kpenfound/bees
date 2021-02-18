package main

type Hive struct {
	Id string
}

func NewHive(id string) *Hive {
	h := &Hive{Id: id}
	return h
}
