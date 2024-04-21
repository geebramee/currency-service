package currencyApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CurrencyClient interface {
	GetCurrencyRate(currencyCode string) (float64, error)
}

type API1Client struct {
	Url string
	Key string
}
type API2Client struct {
	Url string
	Key string
}

type API1Response struct {
	Data map[string]struct {
		Value float64 `json:"value"`
	} `json:"data"`
}

type API2Response struct {
	Rates map[string]float64 `json:"rates"`
}

func (client *API1Client) GetCurrencyRate(currencyCode string) (float64, error) {
	url := fmt.Sprintf(client.Url + client.Key)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var response API1Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	currency, ok := response.Data[currencyCode]
	if !ok {
		return 0, fmt.Errorf("currency not found: %s", currencyCode)
	}

	return currency.Value, nil
}

func (client *API2Client) GetCurrencyRate(currencyCode string) (float64, error) {
	url := fmt.Sprintf(client.Url + client.Key)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var response API2Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	rate, ok := response.Rates[currencyCode]
	if !ok {
		return 0, fmt.Errorf("currency not found: %s", currencyCode)
	}

	return rate, nil
}
