package database

import (
	"github.com/gomodule/redigo/redis"
	"go.uber.org/zap"
)

const (
	ENDPOINT = "redis-15226.c212.ap-south-1-1.ec2.cloud.redislabs.com:15226"
	PASSWORD = "M9XBlksVe8KW7bGYXGS5DwGDPQozPQno"
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
