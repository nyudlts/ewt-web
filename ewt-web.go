package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/nyudlts/ewt-web/lib/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Load configuration from JSON file
	config := router.EWTConfig{}
	configLoc := "ewt-web-config.json"
	configBytes, err := os.ReadFile(configLoc)
	if err != nil {
		log.Fatalf("failed to read config file: %v", err)
	}
	if err := json.Unmarshal(configBytes, &config); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	//initialize the router
	router.InitRouter(r, &config)

	if err := r.Run(); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
