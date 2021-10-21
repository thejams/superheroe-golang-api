package mocks

import "superheroe-api/superheroe-golang-api/src/entity"

var ReqTradeMade = &entity.TradeMade{}

var ResTradeMade = `
{
	"endpoint": "currency data",
	"quotes": [
	 {
	  "ask": 1.15537,
	  "base_currency": "EUR",
	  "bid": 1.15536,
	  "mid": 1.15536,
	  "quote_currency": "USD"
	 },
	 {
	  "ask": 1.3621,
	  "base_currency": "GBP",
	  "bid": 1.36208,
	  "mid": 1.36209,
	  "quote_currency": "USD"
	 }
	],
	"requested_time": "Tue, 12 Oct 2021 11:34:26 GMT",
	"timestamp": 1634038467
 }`

var ResTradeMadeError = `
{
	"message": "internal server error"
}`
