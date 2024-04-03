package main

import (
	"cryptomaster-go/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	print(rate.Price, err)
}
