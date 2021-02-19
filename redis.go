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

}

func (r *Redis) GetHive(id string) Hive {
	return Hive{}
}

func (r *Redis) GetHiveAt(l Location) Hive {
	return Hive{}
}

func (r *Redis) SaveFlower(f Flower) {

}

func (r *Redis) GetFlower(id string) Flower {
	return Flower{}
}

func (r *Redis) GetFlowerAt(l Location) Flower {
	return Flower{}
}

func (r *Redis) SaveBee(b Bee) {

}

func (r *Redis) GetBee(id string) Bee {
	return Bee{}
}

func (r *Redis) See(l Location, distance int) [][]byte {
	// Garbage for now
	s := [][]byte{}
	for i := 0; i < distance*2; i++ {
		for j := 0; j < distance*2; j++ {
			s[i][j] = EmptyCode
		}
	}
	return s
}
