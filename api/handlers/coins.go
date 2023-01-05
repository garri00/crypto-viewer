package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"crypto-viewer/src/config"
	"crypto-viewer/src/entities"
)

type RestyClient struct {
	Usecase CoinsRestyUseCase
}

func NewRestyClient(usecase CoinsRestyUseCase) RestyClient {
	return RestyClient{Usecase: usecase}
}

func (c RestyClient) CoinsResty(w http.ResponseWriter, r *http.Request) {

	restyClient := c.RestyClientAddress

	resp, err := restyClient.R().
		EnableTrace().
		SetQueryParams(map[string]string{
			"start": r.URL.Query().Get("start"),
			"limit": r.URL.Query().Get("limit"),
		}).
		SetHeader("Accepts", "application/json").
		SetHeader("X-CMC_PRO_API_KEY", config.TokenAPI).
		Get("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest")

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET request"))
		return
	}

	fmt.Println(resp.Request.URL)

	if resp.StatusCode() != http.StatusOK {
		var errResponse = entities.Status{}

		if err := json.Unmarshal(resp.Body(), &errResponse); err != nil {
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

	var okResponse = entities.CoinsData{}
	if err := json.Unmarshal(resp.Body(), &okResponse); err != nil {
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
	//fmt.Println(resp.Status)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(file)

}
