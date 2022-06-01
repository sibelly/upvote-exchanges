package service

import (
	"fmt"

	"gopkg.in/resty.v0"
)

// Exchange info
type Exchange struct {
	Exchange_id string `json:"exchange_id"`
	Website     string `json:"website"`
	Name        string `json:"name"`
}

func main() {
	resp, err := resty.R().
		SetHeader("X-CoinAPI-Key", "2E591C1E-3422-41FE-AE88-739CBB128C3C").
		Get("https://rest.coinapi.io/v1/exchanges")

	if err != nil {
		fmt.Println("An error occurred: ", err)
	}

	fmt.Println("Response: ", resp)

}
