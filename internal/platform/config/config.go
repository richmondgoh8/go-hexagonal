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

	log.Printf("reading %s env file\n", environment)
}
