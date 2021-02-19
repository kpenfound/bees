package main

import (
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis() *Redis {
	r := &Redis{}
	r.client = NewRedisClient()
	return r
}

func NewRedisClient() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return rdb
}

func (r *Redis) SaveHive(h Hive) {
	// TODO

}

func (r *Redis) GetHive(id string) Hive {
	// TODO
	return Hive{}
}

func (r *Redis) GetHiveAt(l Location) Hive {
	// TODO
	return Hive{}
}

func (r *Redis) SaveFlower(f Flower) {
	// TODO
}

func (r *Redis) GetFlower(id string) Flower {
	// TODO
	return Flower{}
}

func (r *Redis) DeleteFlower(f *Flower) {
	// TODO
}

func (r *Redis) GetFlowerAt(l Location) Flower {
	// TODO
	return Flower{}
}

func (r *Redis) SaveBee(b Bee) {
	// TODO
}

func (r *Redis) GetBee(id string) Bee {
	// TODO
	return Bee{}
}

func (r *Redis) DeleteBee(b *Bee) {
	// TODO
}

func (r *Redis) GetCodeAt(l Location) byte {
	// TODO
	return EmptyCode
}

func (r *Redis) IsEmpty(l Location) bool {
	return r.GetCodeAt(l) == EmptyCode
}

func (r *Redis) See(l Location, distance int) [][]byte {
	// TODO
	s := [][]byte{}
	for i := 0; i < distance*2; i++ {
		for j := 0; j < distance*2; j++ {
			s[i][j] = EmptyCode
		}
	}
	return s
}
