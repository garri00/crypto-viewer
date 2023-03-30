package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"crypto-viewer/src/entities"
)

var queryParams = map[string]string{
	"start": "1",
	"limit": "4",
}

var okResponse = entities.CoinsData{

	Coins: []entities.Coin{
		{
			ID:     1,
			Name:   "BitCoin",
			Symbol: "BTC",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 20,
				},
			},
		},
		{
			ID:     300,
			Name:   "Etherym",
			Symbol: "ETH",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 10,
				},
			},
		},
		{
			ID:     341,
			Name:   "BNB",
			Symbol: "BNB",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 11255.190172461225,
				},
			},
		},
		{
			ID:     1233,
			Name:   "Tether",
			Symbol: "USDT",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 36.77161508245386,
				},
			},
		},
	},
}

var coinsExchangedUSDtoUAH = entities.CoinsData{

	Coins: []entities.Coin{
		{
			ID:     1,
			Name:   "BitCoin",
			Symbol: "BTC",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 724,
				},
			},
		},
		{
			ID:     300,
			Name:   "Etherym",
			Symbol: "ETH",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 362,
				},
			},
		},
		{
			ID:     341,
			Name:   "BNB",
			Symbol: "BNB",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 407437.8842430964,
				},
			},
		},
		{
			ID:     1233,
			Name:   "Tether",
			Symbol: "USDT",
			Quote: entities.Quote{
				USD: entities.USD{
					Price: 1331.1324659848299,
				},
			},
		},
	},
}

var exchandeRateUSDtoUAH = entities.ExchangeRate{
	Quotes: entities.Quotes{
		USDUAH: 36.2,
	},
	Source:    "USD",
	Success:   true,
	Timestamp: 1675011303,
}

func TestCoinsUseCase_GetCoins(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedError := errors.New("ERROR")

	tests := map[string]struct {
		coinsAdapter    CoinsAdapter
		exchangeAdapter ExchangeAdapter
		exp             entities.CoinsData
		expErr          error
	}{
		"succes": {
			coinsAdapter: func() CoinsAdapter {
				m := NewMockCoinsAdapter(ctrl)
				m.EXPECT().GetCoins(queryParams).Return(okResponse, nil).Times(1)

				return m
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				m := NewMockExchangeAdapter(ctrl)
				m.EXPECT().GetExchangeRate().Return(exchandeRateUSDtoUAH, nil).Times(1)

				return m
			}(),
			exp:    coinsExchangedUSDtoUAH,
			expErr: nil,
		},

		"bad_get_coins": {
			coinsAdapter: func() CoinsAdapter {
				m := NewMockCoinsAdapter(ctrl)
				m.EXPECT().GetCoins(queryParams).Return(entities.CoinsData{}, mockedError).Times(1)

				return m
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				m := NewMockExchangeAdapter(ctrl)
				m.EXPECT().GetExchangeRate().Times(0)

				return m
			}(),
			exp:    entities.CoinsData{},
			expErr: errors.New("cant call coins adapter: " + error(mockedError).Error()),
		},

		"bad_get_exchange_rate": {
			coinsAdapter: func() CoinsAdapter {
				m := NewMockCoinsAdapter(ctrl)
				m.EXPECT().GetCoins(queryParams).Return(okResponse, nil).Times(1)

				return m
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				m := NewMockExchangeAdapter(ctrl)
				m.EXPECT().GetExchangeRate().Return(entities.ExchangeRate{}, mockedError).Times(1)

				return m
			}(),
			exp:    entities.CoinsData{},
			expErr: errors.New("cant call exchange adapter: " + error(mockedError).Error()),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := CoinsUseCase{
				coinsAdapter:    tt.coinsAdapter,
				exchangeAdapter: tt.exchangeAdapter,
			}

			got, err := c.GetCoins(queryParams)
			
			assert.Equal(t, tt.exp, got)
			if tt.expErr != nil {
				assert.EqualError(t, err, tt.expErr.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
