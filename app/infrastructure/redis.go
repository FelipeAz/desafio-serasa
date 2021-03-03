package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/FelipeAz/desafio-serasa/app/interfaces"
	"github.com/garyburd/redigo/redis"
)

// Redis contem a conexao com o REDIS.
type Redis struct {
	Port string
}

// NewRedis retorna um objeto Redis
func NewRedis() interfaces.Redis {
	rds := &Redis{}
	redisPort := fmt.Sprintf("localhost:%s", os.Getenv("REDIS_PORT"))

	rds.Port = redisPort

	return rds
}

// RedisConnect Conecta a porta do Redis e retorna uma instancia da conexao
func (r *Redis) RedisConnect() (redis.Conn, error) {
	c, err := redis.Dial("tcp", r.Port)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// Set insere um valor no Redis identificado pela KEY.
func (r *Redis) Set(key string, value []byte) error {
	conn, err := r.RedisConnect()
	if err != nil {
		log.Println(err)
		return err
	}
	defer conn.Close()

	_, err = conn.Do("SET", key, []byte(value))
	if err != nil {
		log.Println(err)
		return err
	}

	_, err = conn.Do("EXPIRE", key, os.Getenv("REDIS_EXPIRE"))
	if err != nil {
		log.Println(err)
	}

	return err
}

// Get retorna um valor do Redis relacionado a KEY.
func (r *Redis) Get(key string) ([]byte, error) {
	conn, err := r.RedisConnect()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Close()

	var data []byte
	data, err = redis.Bytes(conn.Do("GET", key))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}

// Flush remove um valor do Redis relacionado a KEY.
func (r *Redis) Flush(key string) ([]byte, error) {
	conn, err := r.RedisConnect()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer conn.Close()

	var data []byte
	data, err = redis.Bytes(conn.Do("DEL", key))

	return data, err
}
