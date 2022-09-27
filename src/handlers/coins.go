package handlers

import (
	"crypto-viewer/src/config"
	"crypto-viewer/src/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Coins(w http.ResponseWriter, r *http.Request) {

	u := url.URL{}
	values := u.Query()
	values.Add("start", r.URL.Query().Get("start"))
	values.Add("limit", r.URL.Query().Get("limit"))
	u.RawQuery = values.Encode()

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"+u.String(), nil)
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
			w.Write([]byte("failed to unmarshal errResponse"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Code=%d, Message=%s", errResponse.Status.ErrorCode, errResponse.Status.ErrorMessage)))
		log.Print(errResponse.Status)
		return
	}

	var okResponse = entities.Coins{}

	if err := json.Unmarshal(respBody, &okResponse); err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to unmarshal okResponse"))
		return
	}

	file, err := json.MarshalIndent(okResponse, "", " ")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to marshal okResponse"))
		return
	}

	ioutil.WriteFile("src/pkg/coinslist.json", file, 0644)
	fmt.Println(resp.Status)
	w.WriteHeader(http.StatusOK)

	w.Write(file)

}
