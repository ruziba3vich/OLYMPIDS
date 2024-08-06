package pkg

import (
	"fmt"
	"log"
	configs "medal-service/config"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectPostgres(cnf *configs.DatabaseConfig) (*sqlx.DB, error) {
	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host, cnf.Port, cnf.User, cnf.Password, cnf.Name)

	db, err := sqlx.Connect("postgres", dns)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}

	log.Println("Successfully connected to PostgreSQL")
	return db, nil
}
