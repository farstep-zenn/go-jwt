package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Conn *sql.DB

func init() {
	var err error

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file: ", err)
	}

	user := os.Getenv("DBUser")
	password := os.Getenv("DBPassword")
	host := os.Getenv("DBHost")
	port := os.Getenv("DBPort")
	database := os.Getenv("DBName")

	Conn, err = sql.Open("postgres",
		fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, database))

	if err != nil {
		log.Fatal("OpenError: ", err)
	}

	if err = Conn.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}
}
