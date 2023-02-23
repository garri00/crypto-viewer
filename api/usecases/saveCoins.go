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

func (c SaveDataUseCase) SaveCoins(coinsData entities.CoinsData) error {
	// Write coinsData to file
	file, err := json.MarshalIndent(coinsData, "", " ")
	if err != nil {
		log.Print(err)
		log.Print("failed to unmarshal coinsData")

		return err
	}
	err = ioutil.WriteFile("src/pkg/coinslist.json", file, 0644)
	if err != nil {
		log.Print(err)
		log.Print("failed to save file")

		return err
	}

	return nil
}
