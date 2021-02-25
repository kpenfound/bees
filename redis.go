package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
	ctx    context.Context
}

const (
	RedisLocationPrefix = "location"
	RedisMapKey         = "map"
)

func NewRedis() *Redis {
	r := &Redis{}
	r.client = NewRedisClient()
	r.ctx = context.Background()
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

func (r *Redis) SaveHive(h Hive, setLoc bool) error {
	j, err := json.Marshal(h)
	if err != nil {
		return err
	}

	if setLoc {
		err := r.SetLocation(h.location, HiveCode, h.Id)
		if err != nil {
			return err
		}
	}
	return r.client.Set(r.ctx, h.Id, j, 0).Err()
}

func (r *Redis) GetHive(id string) (Hive, error) {
	j, err := r.client.Get(r.ctx, id).Result()
	if err != nil {
		return Hive{}, err
	}
	var h Hive
	err = json.Unmarshal([]byte(j), &h)
	return h, err
}

func (r *Redis) GetHiveAt(l Location) (Hive, error) {
	h, err := r.GetItemsAt(l)
	if err != nil {
		return Hive{}, err
	}

	for _, hh := range h {
		code, id, _ := breakCoordScoreId(hh)
		if code == HiveCode {
			return r.GetHive(id)
		}
	}

	return Hive{}, nil
}

func (r *Redis) SaveFlower(f Flower, setLoc bool) error {
	j, err := json.Marshal(f)
	if err != nil {
		return err
	}

	if setLoc {
		err := r.SetLocation(f.location, FlowerCode, f.Id)
		if err != nil {
			return err
		}
	}
	return r.client.Set(r.ctx, f.Id, j, 0).Err()
}

func (r *Redis) GetFlower(id string) (Flower, error) {
	j, err := r.client.Get(r.ctx, id).Result()
	if err != nil {
		return Flower{}, err
	}
	var f Flower
	err = json.Unmarshal([]byte(j), &f)
	return f, err
}

func (r *Redis) DeleteFlower(f *Flower) error {
	err := r.DeleteLocation(f.location, FlowerCode, f.Id)
	if err != nil {
		return err
	}
	return r.client.Del(r.ctx, f.Id).Err()
}

func (r *Redis) GetFlowerAt(l Location) (Flower, error) {
	f, err := r.GetItemsAt(l)
	if err != nil {
		return Flower{}, err
	}

	for _, ff := range f {
		code, id, _ := breakCoordScoreId(ff)
		if code == HiveCode {
			return r.GetFlower(id)
		}
	}

	return Flower{}, nil
}

func (r *Redis) SaveBee(b Bee) error {
	j, err := json.Marshal(b)
	if err != nil {
		return err
	}
	err = r.MoveLocation(b.location, BeeCode, b.Id)
	if err != nil {
		return err
	}
	return r.client.Set(r.ctx, b.Id, j, 0).Err()
}

func (r *Redis) GetBee(id string) (Bee, error) {
	j, err := r.client.Get(r.ctx, id).Result()
	if err != nil {
		return Bee{}, err
	}
	var b Bee
	err = json.Unmarshal([]byte(j), &b)
	return b, err
}

func (r *Redis) DeleteBee(b *Bee) error {
	err := r.DeleteLocation(b.location, FlowerCode, b.Id)
	if err != nil {
		return err
	}
	return r.client.Del(r.ctx, b.Id).Err()
}

func redisCoordKey(x int) string {
	return fmt.Sprintf("%s-x-%d", RedisMapKey, x)
}

func redisLocationKey(id string) string {
	return fmt.Sprintf("%s-%s", RedisLocationPrefix, id)
}

func redisCoordScoreKey(code byte, id string, x int) string {
	return fmt.Sprintf("%s:%s:%d", string(code), id, x)
}

func breakCoordScoreId(combinedKey string) (byte, string, int) {
	p := strings.Split(combinedKey, ":")
	code := []byte(p[0])
	id := p[1]
	x, _ := strconv.ParseInt(p[2], 10, 0)
	return code[0], id, int(x)
}

func (r *Redis) GetItemsAt(l Location) ([]string, error) {
	key := redisCoordKey(l.X)
	s, err := r.client.ZRange(r.ctx, key, int64(l.Y), int64(l.Y)).Result()
	return s, err
}

func (r *Redis) MoveLocation(newLoc Location, code byte, id string) error {
	oldLoc, err := r.GetLocation(id)
	if err != nil {
		return err
	}

	err = r.DeleteLocation(oldLoc, code, id)
	if err != nil {
		return err
	}

	err = r.SetLocation(newLoc, code, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) GetLocation(id string) (Location, error) {
	key := redisLocationKey(id)

	s, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		return Location{}, err
	}
	xy := strings.Split(s, ",")
	x, err := strconv.ParseInt(xy[0], 10, 0)
	if err != nil {
		return Location{}, err
	}
	y, err := strconv.ParseInt(xy[1], 10, 0)
	if err != nil {
		return Location{}, err
	}
	return Location{X: int(x), Y: int(y)}, nil
}

func (r *Redis) DeleteLocation(l Location, code byte, id string) error {
	xset := redisCoordKey(l.X)

	scorekey := redisCoordScoreKey(code, id, l.X)

	err := r.client.ZRem(r.ctx, xset, scorekey).Err()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) SetLocation(l Location, code byte, id string) error {
	key := redisLocationKey(id)
	coords := fmt.Sprintf("%d,%d", l.X, l.Y)

	err := r.client.Set(r.ctx, key, coords, 0).Err()
	if err != nil {
		return err
	}

	xset := redisCoordKey(l.X)
	scorekey := redisCoordScoreKey(code, id, l.X)

	err = r.client.ZAdd(r.ctx, xset, &redis.Z{Score: float64(l.Y), Member: scorekey}).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *Redis) See(l Location, distance int, id string) ([][]byte, error) {
	// Empty sight
	sightRangeX := int(math.Min(float64(distance*2), WorldX))
	sightRangeY := int(math.Min(float64(distance*2), WorldY))
	s := make([][]byte, sightRangeX)
	for i := 0; i < sightRangeX; i++ {
		s[i] = make([]byte, sightRangeY)
		for j := 0; j < sightRangeY; j++ {
			s[i][j] = EmptyCode
		}
	}
	// Determine bounds
	startX := int(math.Max(0.0, float64(l.X-distance)))
	startY := int(math.Max(0.0, float64(l.Y-distance)))
	endX := int(math.Min(float64(l.X+distance), WorldX))
	endY := int(math.Min(float64(l.Y+distance), WorldY))

	// Determine sets to query
	var coordSets []string
	for xn := startX; xn <= endX; xn++ {
		coordSets = append(coordSets, redisCoordKey(xn))
	}

	zs := &redis.ZStore{
		Keys:      coordSets,
		Aggregate: "MAX",
	}
	storeKey := fmt.Sprintf("sight-%s", id)

	err := r.client.ZUnionStore(r.ctx, storeKey, zs).Err()
	if err != nil {
		return nil, err
	}

	zr := &redis.ZRangeBy{Min: strconv.Itoa(startY), Max: strconv.Itoa(endY)}
	res, err := r.client.ZRangeByScoreWithScores(r.ctx, storeKey, zr).Result()

	for _, r := range res {
		mem := fmt.Sprintf("%v", r.Member)
		code, _, x := breakCoordScoreId(mem)
		y := int(r.Score)
		s[x][y] = code
	}
	return s, nil
}
