package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConnection contains the database connection string
	StringConnection = ""

	// Port contains the API port
	Port = 0

	// SecretKey contains the JWT secret key
	SecretKey []byte
)

// LoadConfig loads environment variables
func LoadConfig() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	Port, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		Port = 9000
	}

	StringConnection = fmt.Sprintf("%s", os.Getenv("SQLITE_FILE"))
	if err != nil {
		log.Fatal(err)
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
	if err != nil {
		log.Fatal(err)
	}

}
