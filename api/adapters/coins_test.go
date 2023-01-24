package adapters

import (
	"crypto-viewer/src/entities"
	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
	"reflect"
	"testing"
)

func TestCoinsAdapter_GetCoins(t *testing.T) {

	restyClient := resty.New()

	httpmock.ActivateNonDefault(restyClient.GetClient())
	defer httpmock.DeactivateAndReset()
	responder := httpmock.NewStringResponder(200, "")
	httpmock.RegisterResponder("GET", "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?limit=4&start=1", responder)

	tests := map[string]struct {
		name        string
		restyClient *resty.Client
		params      map[string]string
		want        entities.CoinsData // чому це не спрацювало
		wantErr     bool
	}{
		"succes": {
			name:        "succes",
			restyClient: restyClient,
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			want:    entities.CoinsData{},
			wantErr: false,
		},
		"bad responce from CMC api": {
			name:        "bad responce from CMC api",
			restyClient: restyClient,
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			want:    entities.CoinsData{},
			wantErr: false,
		},
		"bad": {
			name:        "bad",
			restyClient: restyClient,
			params: map[string]string{
				"start": "1",
				"limit": "4",
			},
			want:    entities.CoinsData{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CoinsAdapter{
				restyClient: tt.restyClient,
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
