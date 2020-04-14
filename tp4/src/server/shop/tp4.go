package shop

import (
	"log"
	"strconv"
)

type EmptyListError struct {
	s string
}

func (e *EmptyListError) Error() string {
	return e.s
}

type TotalPrice struct {
	Shop  string `json:"tienda"`
	Total int    `json:"precio"`
}

type Prices []TotalPrice

type ProductResult struct {
	result TotalPrice
	err    error
}

func GetTotalPrice(productIds []int, marketNames []string) (Prices, error) {
	if len(productIds) == 0 || len(marketNames) == 0 {
		return nil, &EmptyListError{"Cannot satisfy request with empty lists"}
	}

	var prices = make(Prices, len(marketNames))
	var totalChannel = make(chan ProductResult, len(marketNames))

	for _, marketName := range marketNames {
		go getTotalPriceForMarket(productIds, marketName, totalChannel)
	}

	for i := range prices {
		productResult := <-totalChannel
		if productResult.err != nil {
			return nil, productResult.err
		}
		prices[i] = productResult.result
	}

	return prices, nil
}

func getTotalPriceForMarket(productIds []int, marketName string, totalChannel chan<- ProductResult) {
	var total int
	var productError error
	for _, productId := range productIds {
		stringId := strconv.Itoa(productId)
		product, err := GetProduct(marketName, stringId)
		if err != nil {
			log.Printf("Couldn't get price for product %d, market %s. %s", productId, marketName, err)
			productError = err
			break
		}
		total += product.Price
	}
	totalPrice := TotalPrice{
		Shop:  marketName,
		Total: total,
	}
	totalChannel <- ProductResult{
		result: totalPrice,
		err:    productError,
	}
}
