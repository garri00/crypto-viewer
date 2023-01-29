package usecases

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestCoinsUseCase_GetCoins(t *testing.T) {
	// TODO : change data
	var okResponse = entities.CoinsData{}
	var okResponseBeforeExchange = entities.CoinsData{}
	jsonFile, err := os.Open("test_coinsData_exchanged.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &okResponse)
	json.Unmarshal(byteValue, &okResponseBeforeExchange)

	ctrl := gomock.NewController(t)

	tests := map[string]struct {
		name            string
		coinsAdapter    CoinsAdapter
		exchangeAdapter ExchangeAdapter
		params          map[string]string
		want            entities.CoinsData
		wantErr         bool
	}{
		"succes": {
			name: "succes",
			coinsAdapter: func() CoinsAdapter {
				queryParams := map[string]string{
					"start": "1",
					"limit": "4",
				}
				for i := 0; i < len(okResponseBeforeExchange.Coins); i++ {
					okResponseBeforeExchange.Coins[i].Quote.USD.Price = okResponseBeforeExchange.Coins[i].Quote.USD.Price / 36.2
				}
				m := NewMockCoinsAdapter(ctrl)
				m.EXPECT().GetCoins(queryParams).Return(okResponseBeforeExchange, nil).Times(1)
				return m
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				m := NewMockExchangeAdapter(ctrl)
				m.EXPECT().GetExchangeRate().Return(entities.ExchangeRate{
					Quotes: entities.Quotes{
						USDUAH: 36.2,
					},
					Source:    "USD",
					Success:   true,
					Timestamp: 1675011303,
				},
					nil).
					Times(1)
				return m
			}(),
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			want:    okResponse,
			wantErr: false,
		},

		"bad_get_coins": {
			name: "bad_get_coins",
			coinsAdapter: func() CoinsAdapter {
				queryParams := map[string]string{
					"start": "1",
					"limit": "4",
				}
				m := NewMockCoinsAdapter(ctrl)
				m.EXPECT().GetCoins(queryParams).Return(entities.CoinsData{}, errors.New("cant call coins adapter: ")).Times(1)
				return m
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				m := NewMockExchangeAdapter(ctrl)
				m.EXPECT().GetExchangeRate().Times(0)
				return m
			}(),
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			want:    entities.CoinsData{},
			wantErr: true,
		},

		"bad_get_exchange_rate": {
			name: "bad_get_exchange_rate",
			coinsAdapter: func() CoinsAdapter {
				queryParams := map[string]string{
					"start": "1",
					"limit": "4",
				}
				m := NewMockCoinsAdapter(ctrl)
				m.EXPECT().GetCoins(queryParams).Return(okResponse, nil).Times(1)
				return m
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				m := NewMockExchangeAdapter(ctrl)
				m.EXPECT().GetExchangeRate().Return(entities.ExchangeRate{}, errors.New("cant call exchange adapter: ")).Times(1)
				return m
			}(),
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			want:    entities.CoinsData{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CoinsUseCase{
				coinsAdapter:    tt.coinsAdapter,
				exchangeAdapter: tt.exchangeAdapter,
			}
			got, err := c.GetCoins(tt.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCoins() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCoins() got = %v, want %v", got, tt.want)
			}
		})
	}
}

// Do I need to test this function
//func Test_makeExchange(t *testing.T) {
//	type args struct {
//		coinsData    entities.CoinsData
//		exchangeRate entities.ExchangeRate
//	}
//	tests := []struct {
//		name string
//		args args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			makeExchange(tt.args.coinsData, tt.args.exchangeRate)
//		})
//	}
//}
