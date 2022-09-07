package handlers

import (
	"crypto-viewer/scr/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello")
}

func Coins(w http.ResponseWriter, r *http.Request) {
	u, _ := url.ParseRequestURI("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")
	u.Query().Add("start", r.URL.Query().Get("start"))
	u.Query().Add("limit", r.URL.Query().Get("limit"))
	u.Query().Add("convert", r.URL.Query().Get("convert"))
	u.Query().Add("symbol", r.URL.Query().Get("symbol"))

	fmt.Println(u.EscapedPath())

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com"+u.EscapedPath()+"?start=1&limit=1", nil)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET request"))

		return
	}

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", config.GetConfigTokenAPI())

	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error sending request to server"))
		return
	}

	fmt.Println(resp.Request.URL)
	respBody, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		var errResponse struct {
			Status struct {
				ErrorCode    int    `json:"error_code"`
				ErrorMessage string `json:"error_message"`
			} `json:"status"`
		}

		if err := json.Unmarshal(respBody, &errResponse); err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("failed to unmarshal response"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Code=%d, Message=%s", errResponse.Status.ErrorCode, errResponse.Status.ErrorMessage)))
		log.Print(errResponse.Status)
		return
	}

	fmt.Println(resp.Status)
	w.WriteHeader(http.StatusOK)
	w.Write(respBody)
}
