package shop

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const ProductsPath string = "https://productos-p6pdsjmljq-uc.a.run.app"
const ClientTimeout time.Duration = 10

type Product struct {
	Shop  string `json:"tienda"`
	Id    int    `json:"id"`
	Price int    `json:"precio"`
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{
		Timeout: time.Second * ClientTimeout,
	}
}

type NotFoundError struct {
	s string
}

func (e *NotFoundError) Error() string {
	return e.s
}

func GetProduct(shop string, id string) (*Product, error) {

	productUrl := ProductsPath + "/" + shop + "/productos/" + id

	request, err := http.NewRequest(http.MethodGet, productUrl, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	response, err := Client.Do(request)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer response.Body.Close() // TODO unhandled error

	if response.StatusCode != http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		stringResponse := string(bodyBytes)
		return nil, &NotFoundError{response.Status + stringResponse}
	}

	var decodedResp Product
	decodeError := json.NewDecoder(response.Body).Decode(&decodedResp)
	if decodeError != nil {
		log.Print(decodeError)
		return nil, err
	}
	return &decodedResp, nil
}
