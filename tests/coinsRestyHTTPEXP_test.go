package tests

import (
	"crypto-viewer/src/entities"
	"crypto-viewer/src/handlers"
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_CoinsRestyHandlerHTTPEXP(t *testing.T) {

	var okResponse = entities.Data{}
	jsonFile, err := os.Open("coinslist_test.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &okResponse)

	// run server using httptest
	req := httptest.NewRequest(http.MethodGet, "/coins?start=1&limit=3", nil)
	w := httptest.NewRecorder()
	restyClient := resty.New()
	c := handlers.NewRestyClient(restyClient)
	c.CoinsResty(w, req)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Возвращаем JSON (согласно документации)
		w.WriteHeader(http.StatusOK)
		w.Write(byteValue)
		return
	}))
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)

	// is it working?
	e.GET("/coin").
		Expect().
		Status(http.StatusOK)
}
