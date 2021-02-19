package main

type World struct {
	Hives   []*Hive
	Flowers []*Flower
	redis   *Redis
	nomad   *NomadAPI
}

func NewWorld() *World {
	w := &World{}

	w.nomad = NewNomad()
	w.redis = NewRedis()

	for x := 0; x <= WorldX; x++ {
		for y := 0; y <= WorldY; y++ {
			if x%(WorldX/WorldStartingFlowers) == 0 && y%(WorldY/WorldStartingFlowers) == 0 {
				f := NewFlower(NewId(), Location{X: x, Y: y})
				w.Flowers = append(w.Flowers, f)
			}

			if x%(WorldX/WorldStartingHives) == 0 && y%(WorldY/WorldStartingHives) == 0 {
				h := NewHive(NewId(), Location{X: x, Y: y})
				h.SpawnBees(w.nomad, w.redis, HiveStartingBees)
				w.Hives = append(w.Hives, h)
			}
		}
	}
	return w
}

func (w *World) Simulate() {
	// Do simulation
}

func (w *World) step() {
}

type Location struct {
	X int
	Y int
}

func NewId() string {
	return "a"
}
