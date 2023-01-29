package usecases

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"io/ioutil"
	"log"
)

type SaveDataUseCase struct{}

func NewSaveData() SaveDataUseCase {
	return SaveDataUseCase{}
}

func (c SaveDataUseCase) SaveCoins(coinsData entities.CoinsData) ([]byte, error) {

	//Write coinsData to file
	file, err := json.MarshalIndent(coinsData, "", " ")
	if err != nil {
		log.Print(err)
		log.Print("failed to unmarshal coinsData")
		return nil, err
	}
	ioutil.WriteFile("src/pkg/test_coinsData_exchanged.json", file, 0644)

	return file, nil
}
