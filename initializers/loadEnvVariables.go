package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() (endpoint, key, db, cont string) {
	err := godotenv.Load()
	if err != nil {
		fmt.Print("Errore:", err)
	}

	endpoint = os.Getenv("ENDPOINT")
	key = os.Getenv("COSMOS_API_KEY")
	db = os.Getenv("DATABASE")
	cont = os.Getenv("CONTAINER")

	if endpoint == "" || key == "" || db == "" || cont == "" {
		log.Fatal("Variabili d'ambiente Cosmos mancanti")
	}
	return endpoint, key, db, cont
}
