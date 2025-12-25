package config

import (
	"context"
	"fmt"

	"github.com/Flyingmonk01/go-book-store/pkg/utils"
	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectToDB() *pgx.Conn {
	d, err := pgx.Connect(context.Background(), "postgres://sunny:sameerrai@localhost:5432/postgres?sslmode=disable")
	utils.CheckNilError(err)

	DB = d
	fmt.Println("Connected to database successfully")
	AutoMigrate(DB)
	return DB
}

func AutoMigrate(db *pgx.Conn) {
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS book (
			id VARCHAR(50) PRIMARY KEY,
			name VARCHAR(100) NOT NULL,
			author VARCHAR(100) NOT NULL,
			published_year INT NOT NULL,
			crated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err := db.Exec(context.Background(), createTableSQL)
	utils.CheckNilError(err)
	fmt.Println("Database migrated successfully")
}
