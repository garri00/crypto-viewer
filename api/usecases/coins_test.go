package usecases

import (
	"crypto-viewer/src/entities"
	"reflect"
	"testing"
)

func TestCoinsUseCase_GetCoins(t *testing.T) {
	type args struct {
	}
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
				return nil
			}(),
			exchangeAdapter: func() ExchangeAdapter {
				return nil
			}(),
			params:  nil,
			want:    entities.CoinsData{},
			wantErr: false,
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
