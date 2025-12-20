package config

import (
	"context"
	"fmt"

	"github.com/Flyingmonk01/go-book-store/pkg/utils"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectToDB() {
	d, err := pgx.Connect(context.Background(), "postgres://sunny:sameerrai@localhost:5432/postgres?sslmode=disable")
	utils.CheckNilError(err)

	DB = d
	fmt.Println("Connected to database successfully")
}

func GetDB() *pgx.Conn {
	return DB
}
