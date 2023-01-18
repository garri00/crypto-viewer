package handlers

import (
	"mime/multipart"
	"net/http"
	"net/url"
	"testing"
)

func TestCoinsHandler_CoinsResty(t *testing.T) {

	//restyClient := resty.New()
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

				return nil
			}(),
			saveDataUseCase: func() SaveDataUseCase {

				//queryParams := map[string]string{
				//	"start": "1",
				//	"limit": "4",
				//}

				return nil
			}(),
			response: http.Response{
				Status: http.StatusOK,
			},
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
