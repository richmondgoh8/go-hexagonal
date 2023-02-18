package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func InitReader() {
	environment := ""
	if len(os.Args) < 2 {
		environment = "dev"
	} else {
		environment = os.Args[1]
	}

	err := godotenv.Load(environment + ".env")
	if err != nil {
		log.Fatalf("Error loading %s.env file", environment)
	}

	if os.Getenv("SECRET") == "" {
		log.Fatal("missing SECRET from env file")
	}

	log.Printf("reading %s env file\n", environment)
}
