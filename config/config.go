package config

import (
	"log"
	"os"
)

// var public settings
var (
	Env    string
	Port   string
	APIURL string
)

func init() {

	Env = os.Getenv("ENV")
	Port = os.Getenv("PORT")
	APIURL = os.Getenv("API_URL")

	if Env == "" {
		log.Fatal("cannot find ENV from Env")
	}
	if Port == "" {
		log.Fatal("cannot find PORT from Env")
	}
	if APIURL == "" {
		log.Fatal("cannot find API_URL from Env")
	}
}

//IsProduction to check whether Environment is production
func IsProduction() bool {
	return Env == "production"
}
