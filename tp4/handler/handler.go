package supermarket

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/wildcast/golang-primeros-pasos/tp4/service"

	"github.com/wildcast/golang-primeros-pasos/tp4/model"
)

type caller struct{}

func (c caller) Call(url string) (model.SupermarketResponse, error) {

	timeout := time.Duration(2 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Get(url)

	if err != nil {

		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {

		return model.SupermarketResponse{}, errors.New(strconv.Itoa(res.StatusCode))
	}

	defer res.Body.Close()
	var response model.SupermarketResponse

	json.NewDecoder(res.Body).Decode(&response)

	return response, nil

}

//PriceHandler with will be on charge of perform price operation
func PriceHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	requestData := r.URL.Query()
	url := "https://productos-p6pdsjmljq-uc.a.run.app/name/productos/id"

	supermarketsNames := requestData["supermarkets"]
	productIds := requestData["productids"]
	c := caller{}
	result, error := service.Execute(c, url, supermarketsNames, productIds)

	if error != nil {
		response := model.Response{
			StatusCode: error.Error(),
			Message:    "Problem during client call",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	if result != nil {
		json.NewEncoder(w).Encode(result)
	}

}
