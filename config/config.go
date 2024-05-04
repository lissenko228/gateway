package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	PORT string
)

var Routes = map[string]string{
	"users": "",
}

func GetConf() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	PORT = os.Getenv("PORT")
	Routes["users"] = os.Getenv("USER_SERVICE_URI")
}
