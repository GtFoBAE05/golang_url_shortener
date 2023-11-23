package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DBClient *sqlx.DB

func ConnectPostgres() (db *sqlx.DB, err error) {

	err = godotenv.Load()
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
		panic(fmt.Sprintf("Failed to load .env file: %v", err))
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		return
	}

	DBClient = db
	return
}

func Migrate(db *sqlx.DB) (err error) {
	query := `
		CREATE TABLE IF NOT EXISTS url (
			id SERIAL PRIMARY KEY,
			long_url varchar(150),
			short_url varchar(100),
			access_count int,
			last_accessed timestamp,
			create_at timestamp,
			update_at timestamp
		);
	`

	_, err = db.Exec(query)

	return
}
