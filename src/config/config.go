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
		ClientURI: url,
		Port:      port,
	}
}
