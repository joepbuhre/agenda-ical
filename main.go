package main

import (
	"log"

	"joepbuhre/amphia-agenda-ical/v2/handlers"
	"joepbuhre/amphia-agenda-ical/v2/models"
	"joepbuhre/amphia-agenda-ical/v2/utils"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName(".env") // name of config file (without extension)
	viper.SetConfigType("env")  // REQUIRED if the config file does not have an extension
	viper.AutomaticEnv()
	viper.AddConfigPath(".") // optionally specify the path to look for the config file

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("error reading .env file: %s, this can make sense if you set it manually", err)
	}

	var config = models.Config{
		SecretToken:      viper.GetString("SECRET_TOKEN"),
		DatabaseLocation: viper.GetString("DATABASE_LOCATION"),
	}

	// Initialize the SQLite database
	db, err := utils.InitDB(config.DatabaseLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a Gin router
	r := gin.Default()

	// Serve static files (UI)
	r.Static("/assets", "./assets")

	// Apply authentication middleware and handle UI
	r.GET("/", utils.AuthMiddleware(config.SecretToken), handlers.HandleUI)

	// Initialize API routes with authentication middleware
	api := r.Group("/")
	api.Use(utils.AuthMiddleware(config.SecretToken))

	// Register all agenda routes
	handlers.RegisterAgendaRoutes(api)
	handlers.RegisterMeetingRoutes(api)

	// Handle iCal without middleware
	r.GET("/ical", handlers.HandleIcal)
	r.GET("/ical/json", handlers.HandleJson)

	// Start the server
	log.Println("Server started at :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
