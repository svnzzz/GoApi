package main

import (
	"log"
	"os"
	"strings"
	"tutorial/api/initializers"
	"tutorial/api/routers"
)

func init() {
	if err := initializers.CreateConnection(); err != nil {
		log.Fatal("CreatreConnection error: ", err)
	}

}

func main() {
	r := routers.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	if err := r.Run(port); err != nil {
		log.Fatal(err)
	}
}
