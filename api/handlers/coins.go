package handlers

import (
	"crypto-viewer/api/handlers/usecases"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"crypto-viewer/src/entities"
)

type CoinsHandler struct {
	Usecase usecases.CoinsUsecase
}

func CoinsHendler(usecase usecases.CoinsUsecase) CoinsHandler {
	return CoinsHandler{Usecase: usecase}
}

//func (ch CoinsHendlerContract) CoinsHendlerH()http.Handler{
//	return http.HandlerFunc(func(writer http.ResponseWriter, r *http.Request) {
//		queryParams := map[string]string{
//			"start": r.URL.Query().Get("start"),
//			"limit": r.URL.Query().Get("limit"),
//		}
//
//
//		resp, err := c.Usecase.GetCoinsUC(queryParams)
//	})
//}

func (c CoinsHandler) CoinsResty(w http.ResponseWriter, r *http.Request) {

	queryParams := map[string]string{
		"start": r.URL.Query().Get("start"),
		"limit": r.URL.Query().Get("limit"),
	}

	resp, err := c.Usecase.GetCoinsUC(queryParams)

	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to create GET request"))
		return
	}

	var coinsData = entities.CoinsData{}
	coinsData = resp

	file, err := json.MarshalIndent(coinsData, "", " ")
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed to marshal coinsData"))
		return
	}
	ioutil.WriteFile("src/pkg/coinslist.json", file, 0644)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(file)

}
