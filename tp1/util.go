package tp1

import "strconv"

func convertStringToInt(input string) int {

	conversionResult, error := strconv.Atoi(input)

	if error != nil {
		panic("Problem trying to convert to int")
	}

	return conversionResult
}

func verifyAndProcess(storeName string, price int, storePrices map[string]int) map[string]int {

	_, present := storePrices[storeName]

	if present {

		storePrices[storeName] = storePrices[storeName] + price
	} else {

		storePrices[storeName] = price
	}

	return storePrices

}
