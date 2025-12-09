package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/rillmind/blog/back-end/posts"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	password = "1234"
	user     = "postgres"
	dbname   = "postgres"
)

func Connect() (*gorm.DB, error) {

	dsn := os.Getenv("DATABASE_URL")

	if dsn == "" {
		dsn = fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	if os.Getenv("MIGRATION") == "true" {
		err := db.AutoMigrate(&posts.Model{})
		if err != nil {
			log.Fatal(err)
		}

	}

	// db, err := sql.Open("postgres", dsn)
	// if err != nil {
	// 	panic(err)
	// }

	// err = db.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	fmt.Println("Connected to " + dbname)

	return db, nil
}
