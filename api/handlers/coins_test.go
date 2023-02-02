package handlers

import (
	"crypto-viewer/src/entities"
	"encoding/json"
	"fmt"
	"github.com/gavv/httpexpect/v2"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestCoinsHandler_CoinsResty(t *testing.T) {
	var okResponse = entities.CoinsData{}
	jsonFile, err := os.Open("test_coinsData.json")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &okResponse)

	ctrl := gomock.NewController(t)

	tests := map[string]struct {
		name            string
		coinsUseCase    CoinsUseCase
		saveDataUseCase SaveDataUseCase
		response        http.ResponseWriter
		request         *http.Request
	}{
		"sucess": {
			name: "succes",
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
				file, _ := json.MarshalIndent(okResponse, "", " ")
				m := NewMockSaveDataUseCase(ctrl)
				m.EXPECT().SaveCoins(entities.CoinsData{}).Return(file, nil).Times(1)

				return m
			}(),
			response: nil,
			request:  &http.Request{},
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
		//	response: nil,
		//	request:  &http.Request{},
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
		//	response: nil,
		//	request:  &http.Request{},
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CoinsHandler{
				coinsUseCase:    tt.coinsUseCase,
				saveDataUseCase: tt.saveDataUseCase,
			}
			//

			router := chi.NewRouter()

			c.CoinsResty(tt.response, tt.request)
			//c := CoinsHendler(tt.coinsUseCase, tt.saveDataUseCase)
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

			e.GET("").
				WithQuery("start", "1").WithQuery("limit", "4").
				Expect().
				Status(http.StatusOK)

		})
	}
}
