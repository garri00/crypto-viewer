package usecases

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"io/ioutil"
	"log"
)

type CoinsRestyAdapter interface {
	GetCoinsA(map[string]string) (entities.CoinsData, error)
}

func NewCoinsUC(adapter CoinsRestyAdapter) CoinsUsecase {
	return CoinsUsecase{
		coinsA: adapter,
	}
}

type CoinsUsecase struct {
	coinsA CoinsRestyAdapter
}

func (c CoinsUsecase) GetCoinsUC(params map[string]string) (entities.CoinsData, error) {

	coinsData, err := c.coinsA.GetCoinsA(params)
	if err != nil {
		log.Print(err)
		return coinsData, err
	}

	//Write coinsData to file
	file, err := json.MarshalIndent(coinsData, "", " ")
	if err != nil {
		log.Print(err)
		log.Print("failed to unmarshal coinsData")
		return coinsData, err
	}
	ioutil.WriteFile("src/pkg/coinslist.json", file, 0644)

	//Make exange from USD to UAH

	return coinsData, nil
}

//http request -> handler -> usecase -> adapter -> usecase -> handler -> hhtp response
