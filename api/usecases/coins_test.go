package usecases

import (
	"crypto-viewer/src/entities"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func TestCoinsUseCase_GetCoins(t *testing.T) {

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
				m := NewMockCoinsAdapter(ctrl)
				m.EXPECT().GetCoins(queryParams).Times(1)
				return m
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				m := NewMockExchangeAdapter(ctrl)
				m.EXPECT().GetExchangeRate().Times(1)
				return m
			}(),
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			want:    entities.CoinsData{},
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
				m.EXPECT().GetCoins(queryParams).Times(1)
				return m
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				m := NewMockExchangeAdapter(ctrl)
				m.EXPECT().GetExchangeRate().Times(1).Return(entities.CoinsData{})
				return m
			}(),
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			want:    entities.CoinsData{},
			wantErr: true,
		},
		//
		//"bad_get_exchange_rate": {
		//	name: "succes",
		//	coinsAdapter: func() CoinsAdapter {
		//		queryParams := map[string]string{
		//			"start": "1",
		//			"limit": "4",
		//		}
		//		m := NewMockCoinsAdapter(ctrl)
		//		m.EXPECT().GetCoins(queryParams).Times(1)
		//		return m
		//	}(),
		//	exchangeAdapter: func() ExchangeAdapter {
		//		m := NewMockExchangeAdapter(ctrl)
		//		m.EXPECT().GetExchangeRate().Times(1)
		//		return m
		//	}(),
		//	params: map[string]string{
		//		"start": "1",
		//		"limit": "2",
		//	},
		//	want:    entities.CoinsData{},
		//	wantErr: true,
		//},
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
func Test_makeExchange(t *testing.T) {
	type args struct {
		coinsData    entities.CoinsData
		exchangeRate entities.ExchangeRate
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			makeExchange(tt.args.coinsData, tt.args.exchangeRate)
		})
	}
}
