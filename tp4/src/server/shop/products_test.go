package shop

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	var status int
	var body string
	if strings.Contains(req.URL.Path, "dia") {
		status = http.StatusOK
		body = "{\"tienda\":\"dia\",\"id\":1,\"precio\":7887}"
	} else {
		status = http.StatusBadRequest
		body = "error"
	}
	return &http.Response{
		StatusCode: status,
		Body:       ioutil.NopCloser(bytes.NewBufferString(body)),
	}, nil
}

func TestGetProduct(t *testing.T) {

	oldClient := Client
	defer func() { Client = oldClient }()

	Client = &MockClient{}
	product, err := GetProduct("dia", "1")
	if err != nil {
		t.Errorf("Returns error: %s", err)
	}
	expected := &Product{
		Shop:  "dia",
		Id:    1,
		Price: 7887,
	}
	if !reflect.DeepEqual(product, expected) {
		t.Errorf("Expected %+v, got %+v", expected, product)
	}
	product, err = GetProduct("", "")
	if err == nil {
		t.Errorf("Expecting error")
	}
}
