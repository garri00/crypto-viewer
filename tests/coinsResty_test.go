package tests

import (
	"crypto-viewer/src/entities"
	"crypto-viewer/src/handlers"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func Test_CoinsRestyHandler(t *testing.T) {

	jsonFile, err := os.Open("coinslist_test.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var okResponse = entities.Coins{}
	if err := json.Unmarshal(byteValue, &okResponse); err != nil {
		return
	}

	restyClient := resty.New()

	req, err := http.NewRequest("GET", "https://bd4368da-9638-40bb-b897-2f6ff47d191b.mock.pstmn.io/coins?start=1&limit=3", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(handlers.RestyClientStruct{RestyClientAddress: restyClient}.CoinsResty)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	//expected := okResponse
	//if rr.Body.String() != byteValue {
	//	t.Errorf("handler returned unexpected body: got %v want %v",
	//		rr.Body.String(), expected)
	//}
}
