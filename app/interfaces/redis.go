package interfaces

import "github.com/garyburd/redigo/redis"

// Redis representa a interface do redis.
type Redis interface {
	RedisConnect() (redis.Conn, error)
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Flush(string) ([]byte, error)
}
