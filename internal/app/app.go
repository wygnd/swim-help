package app

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/wygnd/swim-help/internal/config"
)

var conf config.Config

// Invoked before main func
func init() {
	if errLoadDotenv := godotenv.Load(); errLoadDotenv != nil {
		log.Fatal("Error load config...")
	}

	conf = *config.New()
}

// Manage all project package's
func Run() {

}
