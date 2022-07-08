package utils

import (
	"fmt"
	"log"
	"os"

	queries "dexa-training-crud-karyawan/pkg/queries"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Conn *sqlx.DB
}

func DatabaseConnection() (*sqlx.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	// Define database connection for MySQL
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error, not connected to database, %w", err)
	}

	// Try to ping database.
	if err := db.Ping(); err != nil {
		defer db.Close() // close database connection
		return nil, fmt.Errorf("error, not sent ping to database, %w", err)
	}

	return db, nil
}

func DatabaseTest() {
	db, err := DatabaseConnection()
	if err != nil {
		log.Fatal("Failed to Connect to Database: " + err.Error())
	}
	fmt.Println("Database connected!")
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("Database close!")
	}()
}

// Karyawan struct for collect all app queries.
type Karyawan struct {
	*queries.KaryawanQueries // load queries from Book model
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Karyawan, error) {
	// Define a new PostgreSQL connection.
	db, err := DatabaseConnection()
	if err != nil {
		return nil, err
	}

	return &Karyawan{
		// Set queries from models:
		KaryawanQueries: &queries.KaryawanQueries{DB: db}, // from Book model
	}, nil
}
