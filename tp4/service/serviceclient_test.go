package service

import (
	"errors"
	"reflect"
	"testing"

	"github.com/wildcast/golang-primeros-pasos/tp4/model"
)

type mockClient struct {
	fakeResponse model.SupermarketResponse
	message      error
}

func (mock mockClient) Call(url string) (model.SupermarketResponse, error) {

	return mock.fakeResponse, mock.message

}
func TestExecute(t *testing.T) {

	cases := []struct {
		url              string
		names            []string
		ids              []string
		response         model.SupermarketResponse
		customError      error
		expectedResponse map[string]model.Carrito
		expectedError    error
	}{

		{
			url:              "http://name/v1/service",
			names:            nil,
			ids:              nil,
			customError:      errors.New("error"),
			response:         model.SupermarketResponse{},
			expectedError:    nil,
			expectedResponse: map[string]model.Carrito{},
		},
		{
			url:         "http://name/v1/service",
			names:       []string{"Carrefour"},
			ids:         []string{"1"},
			customError: nil,
			response: model.SupermarketResponse{
				Tienda: "Carrefour",
				ID:     1,
				Precio: 100,
			},
			expectedError: nil,
			expectedResponse: map[string]model.Carrito{
				"Carrefour": model.Carrito{
					Precio: 100,
				},
			},
		},
	}

	for _, tc := range cases {

		client := mockClient{
			fakeResponse: tc.response,
			message:      tc.customError,
		}

		got, gotError := Execute(client, tc.url, tc.names, tc.ids)

		if !reflect.DeepEqual(got, tc.expectedResponse) {
			t.Errorf("Got invalid responses")

		}

		if gotError != tc.expectedError {

			t.Errorf("Got '%t' Expected '%t'", gotError, tc.expectedError)
		}

	}

}
