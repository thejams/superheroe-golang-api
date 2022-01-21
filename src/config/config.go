package config

import (
	"os"
	"strings"
	"superheroe-api/superheroe-golang-api/src/entity"
)

//GetAPIConfig returns config struct for the app
func GetAPIConfig() *entity.APPConfig {
	port := os.Getenv("PORT")
	api_key := os.Getenv("API_KEY")
	currencies := os.Getenv("CURRENCIES")
	mongo_usr := os.Getenv("MONGO_USER")
	mongo_pwd := os.Getenv("MONGO_PWD")
	mongo_host := os.Getenv("MONGO_HOST")
	mongo_port := os.Getenv("MONGO_PORT")
	db := os.Getenv("MONGO_DB")

	if len(strings.TrimSpace(mongo_usr)) == 0 {
		mongo_usr = "test"
	}

	if len(strings.TrimSpace(mongo_pwd)) == 0 {
		mongo_pwd = "12345"
	}

	if len(strings.TrimSpace(mongo_host)) == 0 {
		mongo_host = "localhost"
	}

	if len(strings.TrimSpace(mongo_port)) == 0 {
		mongo_port = "27017"
	}

	if len(strings.TrimSpace(db)) == 0 {
		db = "superheroes"
	}

	if len(strings.TrimSpace(port)) == 0 {
		port = ":5000"
	}
	if len(strings.TrimSpace(currencies)) == 0 {
		currencies = "EURUSD,GBPUSD"
	}
	if len(strings.TrimSpace(api_key)) == 0 {
		api_key = "12345"
	}

	url := "https://marketdata.tradermade.com/api/v1/live?currency=" + currencies + "&api_key=" + api_key

	return &entity.APPConfig{
		ClientURI:  url,
		Port:       port,
		MONGO_USER: mongo_usr,
		MONGO_PWD:  mongo_pwd,
		MONGO_HOST: mongo_host,
		MONGO_PORT: mongo_port,
		MONGO_DB:   db,
	}
}
