package utils

import (
	"database/sql"
	"log"
	// "os"
	"time"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	// connStr := os.Getenv("DATABASE_URL")
	connStr := "postgres://admin:admin@localhost:5432/luwa?sslmode=disable"
	var db *sql.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = sql.Open("postgres", connStr)
		if err == nil && db.Ping() == nil {
			break
		}
		log.Println("Waiting for the database to be ready...")
		time.Sleep(2 * time.Second)
	}

	if err != nil || db.Ping() != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	return db
}

func InitDB(db *sql.DB) {
	sqlStatement := `
	CREATE TABLE IF NOT EXISTS users
	(
	    id         SERIAL PRIMARY KEY,
	    email      VARCHAR(255) NOT NULL UNIQUE,
	    password   VARCHAR(255) NOT NULL,
	    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	    packs      JSONB     DEFAULT '[]'
	);

	CREATE TABLE IF NOT EXISTS packs
	(
	    id              SERIAL PRIMARY KEY,
	    brand           VARCHAR(255),
	    model           VARCHAR(255),
	    year_purchased  INT,
	    capacity_litres FLOAT,
	    weight_empty    FLOAT,
	    usage_count     INT
	);

	CREATE TABLE IF NOT EXISTS pack_instances
	(
	    id           SERIAL PRIMARY KEY,
	    pack_id      INT REFERENCES packs (id),
	    total_weight INT,
	    usage_date   TIMESTAMP NOT NULL DEFAULT NOW(),
	    sections     JSONB DEFAULT '[]',
	    user_id      INT REFERENCES users (id),
	    notes        TEXT
	);

	CREATE TABLE IF NOT EXISTS sections
	(
	    id          SERIAL PRIMARY KEY,
	    title       VARCHAR(255),
	    description TEXT,
	    usage_count INT
	);

	CREATE TABLE IF NOT EXISTS section_instances
	(
	    id                        SERIAL PRIMARY KEY,
	    section_id                INT REFERENCES sections (id),
	    pack_instance_id          INT REFERENCES pack_instances (id),
	    total_weight              INT,
	    percentage_of_pack_weight FLOAT,
	    items                    JSONB DEFAULT '[]',
	    notes                    TEXT
	);

	CREATE TABLE IF NOT EXISTS items
	(
	    id          SERIAL PRIMARY KEY,
	    title       VARCHAR(255),
	    description TEXT,
	    weight      FLOAT,
	    usage_count INT
	);

	CREATE TABLE IF NOT EXISTS item_instances
	(
	    id          SERIAL PRIMARY KEY,
	    item_id     INT REFERENCES items (id),
	    section_id  INT REFERENCES sections (id),
	    notes       TEXT
	);
	`

	_, err := db.Exec(sqlStatement)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	log.Println("Database initialized successfully!")
}
