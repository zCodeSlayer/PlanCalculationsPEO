package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"go-postgres/logger"
	"go-postgres/middleware"
	"go-postgres/router"
	"log"
	"net/http"
	"os"
)

func init() {
	if flag.Lookup("test.v") == nil {
		fmt.Println("normal run")
	} else {
		fmt.Println("run under go test")
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error.Println("Error loading .env file")
	}
	port := os.Getenv("GO_POSTGRES_PORT")
	middleware.Port = port
	r := router.Router()
	logger.Info.Println("Starting server on the port " + port + " ...") // TODO: added runserver configurations
	log.Fatal(http.ListenAndServe(":"+port, r))
}
