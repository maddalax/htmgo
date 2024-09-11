package database

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	once sync.Once
	rdb  *redis.Client
)

func Connect() *redis.Client {
	once.Do(func() {
		var ctx = context.Background()
		var err error
		rdb = redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
		if err != nil {
			panic(err)
		}

		cmd := rdb.Ping(ctx)

		if cmd.Err() != nil {
			panic(err)
		}
	})
	return rdb
}

func Incr(key string) int64 {
	db := Connect()
	result := db.Incr(context.Background(), key)
	return result.Val()
}

func Set[T any](key string, value T) error {
	db := Connect()
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}
	result := db.Set(context.Background(), key, serialized, time.Duration(0))
	return result.Err()
}

func HSet[T any](set string, key string, value T) error {
	db := Connect()
	serialized, err := json.Marshal(value)
	if err != nil {
		return err
	}
	result := db.HSet(context.Background(), set, key, serialized)
	return result.Err()
}

func HIncr(set string, key string) int64 {
	db := Connect()
	result := db.HIncrBy(context.Background(), set, key, 1)
	return result.Val()
}

func HGet[T any](set string, key string) *T {
	db := Connect()
	val, err := db.HGet(context.Background(), set, key).Result()
	if err != nil || val == "" {
		return nil
	}
	result := new(T)
	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		return nil
	}
	return result
}

func GetOrSet[T any](key string, cb func() T) (*T, error) {
	db := Connect()
	val, err := db.Get(context.Background(), key).Result()
	if err == nil {
		result := new(T)
		err = json.Unmarshal([]byte(val), result)
		if err != nil {
			return nil, err
		}
		return result, nil
	}
	value := cb()
	err = Set(key, value)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func Get[T any](key string) (*T, error) {
	db := Connect()
	val, err := db.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	result := new(T)
	err = json.Unmarshal([]byte(val), result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func HList[T any](key string) ([]*T, error) {
	db := Connect()
	val, err := db.HGetAll(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	result := make([]*T, len(val))

	count := 0
	for _, t := range val {
		item := new(T)
		err = json.Unmarshal([]byte(t), item)
		if err != nil {
			return nil, err
		}
		result[count] = item
		count++
	}
	return result, nil
}
