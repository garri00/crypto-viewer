package handlers

import (
	"crypto-viewer/src/entities"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"net/http"
	"testing"
	"time"
)

var okResponse = entities.CoinsData{

	Coins: []entities.Coin{
		{
			Id:                0,
			Name:              "",
			Symbol:            "",
			Slug:              "",
			NumMarketPairs:    0,
			DateAdded:         time.Time{},
			MaxSupply:         nil,
			CirculatingSupply: 0,
			TotalSupply:       0,
			Quote: entities.Quote{
				USD: entities.USD{
					Price:                 0,
					Volume24H:             0,
					VolumeChange24H:       0,
					PercentChange1H:       0,
					PercentChange24H:      0,
					PercentChange7D:       0,
					PercentChange30D:      0,
					PercentChange60D:      0,
					PercentChange90D:      0,
					MarketCap:             0,
					MarketCapDominance:    0,
					FullyDilutedMarketCap: 0,
					LastUpdated:           time.Time{},
				},
			},
		},
		{
			Id:                0,
			Name:              "",
			Symbol:            "",
			Slug:              "",
			NumMarketPairs:    0,
			DateAdded:         time.Time{},
			MaxSupply:         nil,
			CirculatingSupply: 0,
			TotalSupply:       0,
			Quote: entities.Quote{
				USD: entities.USD{
					Price:                 0,
					Volume24H:             0,
					VolumeChange24H:       0,
					PercentChange1H:       0,
					PercentChange24H:      0,
					PercentChange7D:       0,
					PercentChange30D:      0,
					PercentChange60D:      0,
					PercentChange90D:      0,
					MarketCap:             0,
					MarketCapDominance:    0,
					FullyDilutedMarketCap: 0,
					LastUpdated:           time.Time{},
				},
			},
		},
		{
			Id:                0,
			Name:              "",
			Symbol:            "",
			Slug:              "",
			NumMarketPairs:    0,
			DateAdded:         time.Time{},
			MaxSupply:         nil,
			CirculatingSupply: 0,
			TotalSupply:       0,
			Quote: entities.Quote{
				USD: entities.USD{
					Price:                 0,
					Volume24H:             0,
					VolumeChange24H:       0,
					PercentChange1H:       0,
					PercentChange24H:      0,
					PercentChange7D:       0,
					PercentChange30D:      0,
					PercentChange60D:      0,
					PercentChange90D:      0,
					MarketCap:             0,
					MarketCapDominance:    0,
					FullyDilutedMarketCap: 0,
					LastUpdated:           time.Time{},
				},
			},
		},
		{
			Id:                0,
			Name:              "",
			Symbol:            "",
			Slug:              "",
			NumMarketPairs:    0,
			DateAdded:         time.Time{},
			MaxSupply:         nil,
			CirculatingSupply: 0,
			TotalSupply:       0,
			Quote: entities.Quote{
				USD: entities.USD{
					Price:                 0,
					Volume24H:             0,
					VolumeChange24H:       0,
					PercentChange1H:       0,
					PercentChange24H:      0,
					PercentChange7D:       0,
					PercentChange30D:      0,
					PercentChange60D:      0,
					PercentChange90D:      0,
					MarketCap:             0,
					MarketCapDominance:    0,
					FullyDilutedMarketCap: 0,
					LastUpdated:           time.Time{},
				},
			},
		},
	},
}

func TestCoinsHandler_CoinsResty(t *testing.T) {
	ctrl := gomock.NewController(t)

	tests := map[string]struct {
		coinsUseCase    CoinsUseCase
		saveDataUseCase SaveDataUseCase
	}{
		"sucess": {
			coinsUseCase: func() CoinsUseCase {
				queryParams := map[string]string{
					"start": "1",
					"limit": "4",
				}
				m := NewMockCoinsUseCase(ctrl)
				m.EXPECT().GetCoins(queryParams).Return(okResponse, nil).Times(1)

				return m
			}(),
			saveDataUseCase: func() SaveDataUseCase {
				m := NewMockSaveDataUseCase(ctrl)
				m.EXPECT().SaveCoins(okResponse).Return(nil).Times(1)

				return m
			}(),
		},

		//"bad coins usecase": {
		//	name: "bad coins usecase",
		//	coinsUseCase: func() CoinsUseCase {
		//		queryParams := map[string]string{
		//			"start": "1",
		//			"limit": "4",
		//		}
		//		m := NewMockCoinsUseCase(ctrl)
		//		m.EXPECT().GetCoins(queryParams).Times(1)
		//
		//		return m
		//	}(),
		//	saveDataUseCase: func() SaveDataUseCase {
		//		m := NewMockSaveDataUseCase(ctrl)
		//		m.EXPECT().SaveCoins(entities.CoinsData{}).Times(1)
		//
		//		return m
		//	}(),
		//},

		//"bad save data usecase": {
		//	name: "bad save data usecase",
		//	coinsUseCase: func() CoinsUseCase {
		//		queryParams := map[string]string{
		//			"start": "1",
		//			"limit": "4",
		//		}
		//		m := NewMockCoinsUseCase(ctrl)
		//		m.EXPECT().GetCoins(queryParams).Times(1)
		//
		//		return m
		//	}(),
		//	saveDataUseCase: func() SaveDataUseCase {
		//		m := NewMockSaveDataUseCase(ctrl)
		//		m.EXPECT().SaveCoins(entities.CoinsData{}).Times(1)
		//
		//		return m
		//	}(),
		//},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := CoinsHandler{
				coinsUseCase:    tt.coinsUseCase,
				saveDataUseCase: tt.saveDataUseCase,
			}

			router := chi.NewRouter()
			router.Get("/coins", c.CoinsResty)
			httpClient := &http.Client{
				Transport: httpexpect.NewBinder(router),
				Jar:       httpexpect.NewJar(),
			}

			e := httpexpect.WithConfig(httpexpect.Config{
				Client:   httpClient,
				Reporter: httpexpect.NewAssertReporter(t),
				Printers: []httpexpect.Printer{
					httpexpect.NewDebugPrinter(t, true),
				},
			})

			e.GET("/coins").
				WithQuery("start", "1").WithQuery("limit", "4").
				Expect().
				Status(http.StatusOK).Body()

		})
	}
}
