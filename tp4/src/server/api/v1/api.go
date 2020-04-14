package v1

import (
	"encoding/json"
	"log"
	"net/http"
	"server/shop"
	"strconv"

	"github.com/go-chi/chi"
)

func pricesResponse(writer http.ResponseWriter, prices shop.Prices) {
	err := json.NewEncoder(writer).Encode(prices)
	if err != nil {
		http.Error(writer, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func GetApiHandler() http.Handler {
	router := chi.NewRouter()
	router.Get("/prices", getPrices)
	return router
}

func getPrices(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()

	productIds := query["products"]
	products := make([]int, len(productIds))
	for i, productId := range productIds {
		product, err := strconv.Atoi(productId)
		if err != nil {
			log.Printf("%s is not a valid product id", productId)
		} else {
			products[i] = product
		}
	}
	prices, err := shop.GetTotalPrice(products, query["markets"])
	if err != nil {
		handleError(err, writer)
	}
	pricesResponse(writer, prices)
}

func handleError(err error, writer http.ResponseWriter) {
	switch e := err.(type) {
	case *shop.NotFoundError:
		http.Error(writer, e.Error(), http.StatusNotFound)
	case *shop.EmptyListError:
		http.Error(writer, e.Error(), http.StatusBadRequest)
	default:
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
