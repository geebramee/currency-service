package main

import (
	"currency-service/cache"
	"currency-service/config"
	service "currency-service/services"
	"currency-service/transport/currencyApi"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	cfg := config.GetConfig()

	c := cache.NewCache(cfg.RedisUrl, cfg.RedisPassword)

	cs := &service.CurrencyService{
		Config: cfg,
		Cache:  c,
		API1Client: &currencyApi.API1Client{
			Url: "https://api.currencyapi.com/v3/latest",
			Key: "?apikey=cur_live_3Jb5MzRPndFvRrKSD1VsrPnvjLxgh1OzzaHp64nL",
		},
		API2Client: &currencyApi.API2Client{
			Url: "https://api.currencyfreaks.com/v2.0/rates/latest",
			Key: "?apikey=b4797efc22354bd399187f1caa03c344",
		},
	}

	http.HandleFunc("/currency/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) < 3 {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		currencyCode := parts[2]

		rate, err := cs.GetCurrencyRate(currencyCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(map[string]interface{}{"rate": rate})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	})

	http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), nil)
}
