package config

import (
	"os"
	"strings"
)

type APPConfig struct {
	TradeMadeClientURI string
	Port               string
	MONGO_USER         string
	MONGO_PWD          string
	MONGO_HOST         string
	MONGO_PORT         string
	MONGO_DB           string
}

//GetAPIConfig returns config struct for the app
func GetAPIConfig() *APPConfig {
	port := os.Getenv("PORT")

	if len(strings.TrimSpace(port)) == 0 {
		port = ":5000"
	}

	mongoSetting := getMongoConfig()

	return &APPConfig{
		TradeMadeClientURI: getTradeMadeConfig(),
		Port:               port,
		MONGO_USER:         mongoSetting["MONGO_USER"],
		MONGO_PWD:          mongoSetting["MONGO_PWD"],
		MONGO_HOST:         mongoSetting["MONGO_HOST"],
		MONGO_PORT:         mongoSetting["MONGO_PORT"],
		MONGO_DB:           mongoSetting["MONGO_DB"],
	}
}

// getTradeMadeConfig sets the trade made client settings
func getTradeMadeConfig() string {
	trade_made_api_key := os.Getenv("API_KEY")
	trade_made_currencies := os.Getenv("CURRENCIES")

	if len(strings.TrimSpace(trade_made_currencies)) == 0 {
		trade_made_currencies = "EURUSD,GBPUSD"
	}
	if len(strings.TrimSpace(trade_made_api_key)) == 0 {
		trade_made_api_key = "12345"
	}

	return "https://marketdata.tradermade.com/api/v1/live?currency=" + trade_made_currencies + "&api_key=" + trade_made_api_key
}

// getMongoConfig sets the mongo db settings
func getMongoConfig() map[string]string {
	m := make(map[string]string)
	m["MONGO_USER"] = os.Getenv("MONGO_USER")
	m["MONGO_PWD"] = os.Getenv("MONGO_PWD")
	m["MONGO_HOST"] = os.Getenv("MONGO_HOST")
	m["MONGO_PORT"] = os.Getenv("MONGO_PORT")
	m["MONGO_DB"] = os.Getenv("MONGO_DB")

	if len(strings.TrimSpace(m["MONGO_USER"])) == 0 {
		m["MONGO_USER"] = "test"
	}

	if len(strings.TrimSpace(m["MONGO_PWD"])) == 0 {
		m["MONGO_PWD"] = "12345"
	}

	if len(strings.TrimSpace(m["MONGO_HOST"])) == 0 {
		m["MONGO_HOST"] = "localhost"
	}

	if len(strings.TrimSpace(m["MONGO_PORT"])) == 0 {
		m["MONGO_PORT"] = "27017"
	}

	if len(strings.TrimSpace(m["MONGO_DB"])) == 0 {
		m["MONGO_DB"] = "superheroes"
	}

	return m
}
