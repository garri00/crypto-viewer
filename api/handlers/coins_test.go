package handlers

import (
	"crypto-viewer/src/entities"
	"github.com/golang/mock/gomock"
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"
)

func TestCoinsHandler_CoinsResty(t *testing.T) {

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
				m.GetCoins(queryParams)

				return m
			}(),
			saveDataUseCase: func() SaveDataUseCase {
				m := NewMockSaveDataUseCase(ctrl)
				m.SaveCoins(entities.CoinsData{})

				return m
			}(),
			response: nil,
			request: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Scheme:      "",
					Opaque:      "",
					User:        &url.Userinfo{},
					Host:        "",
					Path:        "",
					RawPath:     "",
					OmitHost:    false,
					ForceQuery:  false,
					RawQuery:    "",
					Fragment:    "",
					RawFragment: "",
				},
				Proto:            "",
				ProtoMajor:       0,
				ProtoMinor:       0,
				Header:           nil,
				Body:             nil,
				GetBody:          nil,
				ContentLength:    0,
				TransferEncoding: nil,
				Close:            false,
				Host:             "",
				Form:             nil,
				PostForm:         nil,
				MultipartForm: &multipart.Form{
					Value: nil,
					File:  nil,
				},
				Trailer:    nil,
				RemoteAddr: "",
				RequestURI: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := CoinsHandler{
				coinsUseCase:    tt.coinsUseCase,
				saveDataUseCase: tt.saveDataUseCase,
			}
			c.CoinsResty(tt.response, tt.request)
		})
	}
}
