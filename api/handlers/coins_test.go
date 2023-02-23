package handlers

import (
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"

	"crypto-viewer/src/entities"
)

var okResponse = entities.CoinsData{

	Coins: []entities.Coin{
		{
			ID:                0,
			Name:              "BTC",
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
			ID:                0,
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
			ID:                0,
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
			ID:                0,
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
		expQuery        map[string]string
		expStatus       int
		expBody         string
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
			expQuery: map[string]string{
				"start": "1",
				"limit": "4",
			},
			expStatus: http.StatusOK,
			// how to send exp body in succes
			expBody: "\"name\": \"BTC\"",
		},

		"bad query params": {
			coinsUseCase: func() CoinsUseCase {
				queryParams := map[string]string{
					"start": "0",
					"limit": "4",
				}
				m := NewMockCoinsUseCase(ctrl)
				m.EXPECT().GetCoins(queryParams).Times(0)

				return m
			}(),
			saveDataUseCase: func() SaveDataUseCase {
				m := NewMockSaveDataUseCase(ctrl)
				m.EXPECT().SaveCoins(nil).Times(0)

				return m
			}(),
			expQuery: map[string]string{
				"start": "0",
				"limit": "4",
			},
			expStatus: http.StatusInternalServerError,
			expBody:   "wrong query pqrams",
		},

		"bad coins usecase": {
			coinsUseCase: func() CoinsUseCase {
				queryParams := map[string]string{
					"start": "1",
					"limit": "4",
				}
				m := NewMockCoinsUseCase(ctrl)
				m.EXPECT().GetCoins(queryParams).Return(entities.CoinsData{}, errors.New("cant call coins adapter:")).Times(1)

				return m
			}(),
			saveDataUseCase: func() SaveDataUseCase {
				m := NewMockSaveDataUseCase(ctrl)
				m.EXPECT().SaveCoins(nil).Times(0)

				return m
			}(),
			expQuery: map[string]string{
				"start": "1",
				"limit": "4",
			},
			expStatus: http.StatusInternalServerError,
			expBody:   "failed to create GET coins",
		},

		"bad save data usecase": {
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
				m.EXPECT().SaveCoins(okResponse).Return(errors.New("failed to unmarshal coinsData")).Times(1)

				return m
			}(),
			expQuery: map[string]string{
				"start": "1",
				"limit": "4",
			},
			expStatus: http.StatusInternalServerError,
			expBody:   "failed to save coins",
		},
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
			}

			e := httpexpect.WithConfig(httpexpect.Config{
				Client:   httpClient,
				Reporter: httpexpect.NewAssertReporter(t),
				Printers: []httpexpect.Printer{
					httpexpect.NewDebugPrinter(t, true),
				},
			})

			e.GET("/coins").
				// do I need to chech query params here
				WithQuery("start", tt.expQuery["start"]).WithQuery("limit", tt.expQuery["limit"]).
				Expect().
				Status(tt.expStatus).
				Body().Contains(tt.expBody)

		})
	}
}
