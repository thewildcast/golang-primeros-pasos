package client

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"github.com/wildcast/golang-primeros-pasos/tp4/model"
)

func Call(url string, w http.ResponseWriter) model.SupermarketResponse {

	timeout := time.Duration(2 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}

	res, err := client.Get(url)

	if err != nil {

		log.Fatal(err)
	}

	if res.StatusCode != http.StatusOK {

		w.WriteHeader(http.StatusNotFound)
		error := model.Response {
			Error: "Problem during client call",
			StatusCode : res.StatusCode,
		}
		json.NewEncoder(w).Encode(error)

		return model.SupermarketResponse{
			Tienda: "",
			Precio: 0,
		}

	}

	defer res.Body.Close()
	var response model.SupermarketResponse

	json.NewDecoder(res.Body).Decode(&response)

	return response

}