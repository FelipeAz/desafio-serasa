package main

import (
	"log"

	"github.com/FelipeAz/desafio-serasa/config/infrastructure"
)

func main() {
	db, err := infrastructure.NewSQLHandler()
	if err != nil {
		log.Fatal(err)
	}
	defer db.CloseConnection()

	rds := infrastructure.NewRedis()
	testConn, err := rds.RedisConnect()
	if err != nil {
		log.Fatal(err)
	}
	testConn.Close()

	router := infrastructure.NewRouter()
	router.Dispatch(db, rds)
}
