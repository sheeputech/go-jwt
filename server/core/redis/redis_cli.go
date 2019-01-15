package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

type RedisCli struct {
	conn redis.Conn
}

var instanceRedisCli *RedisCli = nil

func Connect() (conn *RedisCli) {
	if instanceRedisCli == nil {
		instanceRedisCli = new(RedisCli)
		var err error

		instanceRedisCli.conn, err = redis.Dial("tcp", "192.168.99.100:6379")
		if err != nil {
			fmt.Println("connection failed: ", err)
			panic(err)
		}

		if _, err := instanceRedisCli.conn.Do("AUTH", "SheepuTech"); err != nil {
			instanceRedisCli.conn.Close()
			fmt.Println("authorization failed: ", err)
			panic(err)
		}
	}
	return instanceRedisCli
}

func (redisCli *RedisCli) SetValue(key string, value string, expiration ...interface{}) error {
	_, err := redisCli.conn.Do("SET", key, value)

	if err != nil && expiration != nil {
		redisCli.conn.Do("EXPIRE", key, expiration[0])
	}

	return err
}

func (redisCli *RedisCli) GetValue(key string) (interface{}, error) {
	return redisCli.conn.Do("GET", key)
}
