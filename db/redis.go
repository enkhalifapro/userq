package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

// NewPool creates a connection pool.
func NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", viper.GetString("db.uri"))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
