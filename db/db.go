package db

import (
	"github.com/go-pg/pg/v10"
	_ "github.com/go-pg/pg/v10/orm"
)


func ConnectDB() *pg.DB {
	opt, err := pg.ParseURL("postgres://postgres:postgres@localhost:15432/postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	return db
}

