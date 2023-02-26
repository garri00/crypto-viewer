package usecases

import (
	"encoding/json"
	"os"

	"github.com/rs/zerolog"

	"crypto-viewer/src/entities"
)

type SaveDataUseCase struct {
	log zerolog.Logger
}

func NewSaveData(l zerolog.Logger) SaveDataUseCase {
	return SaveDataUseCase{
		log: l,
	}
}

func (c SaveDataUseCase) SaveCoins(coinsData entities.CoinsData) error {
	// Write coinsData to file
	file, err := json.MarshalIndent(coinsData, "", " ")
	if err != nil {
		c.log.Error().Err(err).Msg("failed to unmarshal coinsData")

		return err
	}

	perm := 0600
	err = os.WriteFile("src/pkg/coinslist.json", file, os.FileMode(perm))
	if err != nil {
		c.log.Error().Err(err).Msg("failed to save file")

		return err
	}

	return nil
}
