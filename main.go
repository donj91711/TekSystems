package main

import (
	"TEKSystems/database"
	"TEKSystems/handler"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Config struct {
	DBUser     string `json:"dbUser"`
	DBPassword string `json:"dbPassword"`
	DBHost     string `json:"dbHost"`
	DBName     string `json:"dbName"`
	DBPort     string `json:"dbPort"`
	HTTPPort   string `json:"httpPort"`
}

func main() {
	fmt.Println("Starting API Implementation Exercise")
	config := getConfig(0) //0 for Docker built DB, 1 for localhost manually built db

	db, err := database.OpenDB(config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	router := mux.NewRouter()
	handler.SetRoutes(router, *db)
	go func() {
		err := http.ListenAndServe(":"+config.HTTPPort, router)
		if err != nil {
			fmt.Println("Error starting HTTP server:", err)
		}
	}()

	// Keep the main goroutine alive
	fmt.Println("*** TekSystems Application is now running ***")
	select {}
}

func getConfig(opt int) Config {
	configFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return Config{}
	}
	defer configFile.Close()

	var configWrapper struct {
		Config []Config `json:"config"`
	}

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&configWrapper)
	if err != nil {
		fmt.Println("Error decoding config file:", err)
		return Config{}
	}

	if opt == 0 {
		return configWrapper.Config[0] //running inside a Docker container
	} else {
		return configWrapper.Config[1] //running locally
	}
}
