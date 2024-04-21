package currencyApi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPI1Client_GetCurrencyRate(t *testing.T) {
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mockResponse := API1Response{
			Data: map[string]struct {
				Value float64 `json:"value"`
			}{
				"USD": {Value: 1.2},
				"EUR": {Value: 1.3},
			},
		}

		responseBody, _ := json.Marshal(mockResponse)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	})

	mockServer := httptest.NewServer(mockHandler)
	defer mockServer.Close()

	apiClient := &API1Client{
		Url: mockServer.URL,
		Key: "?apikey=testkey",
	}

	testCases := []struct {
		CurrencyCode     string
		ExpectedRate     float64
		ExpectedError    error
		ExpectedErrorMsg string
	}{
		{"USD", 1.2, nil, ""},
		{"EUR", 1.3, nil, ""},
		{"GBP", 0, fmt.Errorf("currency not found: GBP"), "currency not found: GBP"},
	}

	for _, tc := range testCases {
		rate, err := apiClient.GetCurrencyRate(tc.CurrencyCode)

		if rate != tc.ExpectedRate {
			t.Errorf("for currency code %s, expected rate %f, got %f", tc.CurrencyCode, tc.ExpectedRate, rate)
		}

		if err != nil {
			if err.Error() != tc.ExpectedErrorMsg {
				t.Errorf("for currency code %s, expected error '%s', got '%s'", tc.CurrencyCode, tc.ExpectedErrorMsg, err.Error())
			}
		} else {
			if tc.ExpectedError != nil {
				t.Errorf("for currency code %s, expected error '%s', got nil", tc.CurrencyCode, tc.ExpectedError.Error())
			}
		}
	}
}
