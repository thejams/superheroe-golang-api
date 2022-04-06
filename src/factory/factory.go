package factory

import (
	"superheroe-api/superheroe-golang-api/src/client"
	"superheroe-api/superheroe-golang-api/src/repository"
)

func DBFactory(t int) repository.Repository {
	switch t {
	case 1:
		return &repository.MongoDB{}
	default:
		return nil
	}
}

func ClientFactory(t int) client.Client {
	switch t {
	case 1:
		return &client.TradeMadeClient{}
	default:
		return nil
	}
}
