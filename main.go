package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db *sql.DB

	dbHost     string
	dbUsername string
	dbName     string
	dbPort     string
	dbPassword string
)

func init() {
	//Load .env file
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file.")
	}

	dbHost = os.Getenv("dbHost")
	dbUsername = os.Getenv("dbUsername")
	dbPassword = os.Getenv("dbPassword")
	dbName = os.Getenv("dbName")
	dbPort = os.Getenv("dbPort")
}

func main() {
	var err error

	pgConString := fmt.Sprintf("port=%s host=%s user=%s "+"password=%s dbname=%s sslmode=disable", dbPort, dbHost, dbUsername, dbPassword, dbName)
	db, err := sql.Open("postgres", pgConString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	//Flags initialize BD.
	initDb := flag.Bool("init", true, "Initialize tables in DB.")
	seed := flag.Bool("seed", true, "Fill 10 years information in BD.")
	flag.Parse()

	if *initDb {
		if err := database.initializeDb(db); err != nil {
			log.Fatalf("Error initializing database: %v\n", err)
		}
	}

	if *seed {
		if err := database.seedDb(db); err != nil {
			log.Fatalf("Error seeding database: %v\n", err)
		}
	}
}
