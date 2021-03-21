package database

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)
// Add you endpoint and pass for Redis Cloud
const (
	ENDPOINT = "endpoint"
	PASSWORD = "password"
)

func ConnectRedis() (redis.Conn, error) {

	// Setup connection to Redis Cloud
	conn, err := redis.Dial("tcp", ENDPOINT, redis.DialPassword(PASSWORD))

	if err != nil {
		// fmt.Println(err.Error())
		zap.String("Error: Redis", err.Error())
	}

	return conn, err
}
