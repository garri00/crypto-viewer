package usecases

import (
	"crypto-viewer/api/handlers/adapters"
	"crypto-viewer/src/entities"
	"net/http"
)

type CoinsRestyUseCase interface {
	GetCoins(http.Response) (entities.CoinsData, error)
}

type Coins struct {
	coinsUC CoinsRestyUseCase
	coinsA  adapters.CoinsRestyAdapter
}

//func (c Coins) GetCoins(resp http.Response) (entities.CoinsData, error) {
//
//	if resp.StatusCode != http.StatusOK {
//		var errResponse = entities.Status{}
//
//		if err := json.Unmarshal(resp.Request.Body(), &errResponse); err != nil {
//			log.Print(err)
//			return nil, nil
//		}
//
//		w.WriteHeader(http.StatusInternalServerError)
//		w.Write([]byte(fmt.Sprintf("Code=%d, Message=%s", errResponse.Status.ErrorCode, errResponse.Status.ErrorMessage)))
//		log.Print(errResponse.Status)
//		return
//	}
//	return coins, nil
//
//}
