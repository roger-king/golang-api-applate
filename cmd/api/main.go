package main

import (
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload" // blank
	server "{{initVar}}/pkg"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	// Set logging settings
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	server, _ := server.New()
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "9000"
	}
	
	log.Fatal(server.Listen(fmt.Sprintf(":%s", port)))
}
