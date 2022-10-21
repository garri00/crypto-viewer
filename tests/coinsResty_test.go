package tests

import (
	"crypto-viewer/src/handlers"
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CoinsRestyHandler(t *testing.T) {
	//server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	// повертаємо дані
	//	w.WriteHeader(http.StatusOK)
	//	w.Write([]byte(`{"data":{"id":"bitcoin","symbol":"BTC","currencySymbol":"₿","type":"crypto","rateUsd":"4010.8714336221081818"},"timestamp":1552990697033}`))
	//	return
	//}))

	//restyClient := resty.New()
	//resources := handlers.RestyClientStruct{RestyClientAddress: restyClient}

	//resources.CoinsResty()

	//if err != nil {
	//	t.Error(err)
	//}
	////
	//if result == 0 {
	//	t.Fail()
	//}
	//expected := "dummy data"
	//svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	w.WriteHeader(http.StatusOK)
	//	fmt.Fprintf(w, expected)
	//}))
	//defer svr.Close()
	//c := NewClient(svr.URL)
	//
	//if err != nil {
	//	t.Errorf("expected err to be nil got %v", err)
	//}
	//// res: expected\r\n
	//// due to the http protocol cleanup response
	//res = strings.TrimSpace(res)
	//if res != expected {
	//	t.Errorf("expected res to be %s got %s", expected, res)
	//}

	req := httptest.NewRequest(http.MethodGet, "/coins?start=1&limit=3", nil)
	w := httptest.NewRecorder()

	restyClient := resty.New()
	c := handlers.NewRestyClient(restyClient)
	c.CoinsResty(w, req)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code not 200")
	}
}
