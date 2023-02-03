package client_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"superheroe-api/superheroe-golang-api/src/client"
	"superheroe-api/superheroe-golang-api/src/client/mocks"
	"superheroe-api/superheroe-golang-api/src/entity"
)

func Test_client_trade_made(t *testing.T) {
	t.Run("ok response trade_made_client", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(mocks.ResTradeMade))
		}))
		defer server.Close()

		tradeMadeClient := client.NewTradeMadeClient(server.URL)
		res, err := tradeMadeClient.Get()
		tradeMadeObj, ok := res.(entity.TradeMade)

		assert.Nil(t, err)
		assert.True(t, ok)
		assert.Equal(t, tradeMadeObj.Endpoint, "currency data")

	})

	t.Run("fail response with 500 error trade_made_client", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(mocks.ResTradeMadeError))
		}))
		defer server.Close()

		tradeMadeClient := client.NewTradeMadeClient(server.URL)
		_, err := tradeMadeClient.Get()

		assert.NotNil(t, err)
	})

	t.Run("fail response with 404 error trade_made_client", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(mocks.ResTradeMadeError))
		}))
		defer server.Close()

		tradeMadeClient := client.NewTradeMadeClient(server.URL)

		_, err := tradeMadeClient.Get()
		assert.NotNil(t, err)
	})
}
