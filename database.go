package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

var DB *pgxpool.Pool

var err error

func ConnectDB() {

	// urlExample := "postgres://username:password@localhost:5432/database_name"
	DB, err = pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		"postgres",
		"1234",
		"127.0.0.1",
		"5432",
		"camp"))
	if err != nil {
		fmt.Print(err)
	}

}
