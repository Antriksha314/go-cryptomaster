package api

import (
	"cryptomaster-go/model"
	"fmt"
	"io"
	"net/http"
	"strings"
)

const apiURL = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*model.Rate, error) {
	upperCurrency := strings.ToUpper(currency)

	res, err := http.Get(fmt.Sprintf(apiURL, upperCurrency))
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		data, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		json := string(data)
		fmt.Println(json)
	} else {
		return nil, fmt.Errorf("status code returned: %v", res.StatusCode)
	}
	rate := model.Rate{Currency: upperCurrency, Price: 1.0}
	return &rate, nil
}
