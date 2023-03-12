package main

import (
	"app/router"
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDsn() string {
	host := "host=" + os.Getenv("POSTGRES_HOST")
	user := "user=" + os.Getenv("POSTGRES_USER")
	password := "password=" + os.Getenv("POSTGRES_PASSWORD")
	port := "port=" + os.Getenv("POSTGRES_PORT")
	dbname := "dbname=" + os.Getenv("POSTGRES_DBNAME")
	sslmode := "sslmode=" + os.Getenv("POSTGRES_SSLMODE")
	timezone := "TimeZone=" + os.Getenv("POSTGRES_TIMEZONE")

	return strings.Join([]string{host, user, password, port, dbname, sslmode, timezone}, " ")
}

func dbConnect() *gorm.DB {
	dsn := getDsn()
	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if dbErr != nil {
		log.Fatal(dbErr)
	}

	return db
}

func main() {
	dbConnect()
	engine := router.Load()

	engine.Run()
}
