package service

import (
	"currency-service/cache"
	"currency-service/config"
	"currency-service/transport/currencyApi"
	"sync"
)

type CurrencyServiceI interface {
	GetCurrencyRate(currencyCode string) (float64, error)
}

type CurrencyService struct {
	Config     *config.Config
	Cache      *cache.Cache
	API1Client currencyApi.CurrencyClient
	API2Client currencyApi.CurrencyClient
	Mutex      sync.Mutex
	Count      int
}

func (cs *CurrencyService) GetCurrencyRate(currencyCode string) (float64, error) {
	cs.Mutex.Lock()
	defer cs.Mutex.Unlock()

	rate, err := cs.Cache.Get(currencyCode)
	if err == nil {
		return rate, nil
	}

	cs.Count++
	useAPI1 := (cs.Count % 100) < int(cs.Config.API1Percent)
	if useAPI1 {
		rate, err = cs.API1Client.GetCurrencyRate(currencyCode)
	} else {
		rate, err = cs.API2Client.GetCurrencyRate(currencyCode)
	}

	if err != nil {
		rate, cacheErr := cs.Cache.Get(currencyCode)
		if cacheErr == nil {
			return rate, nil
		}
		return 0, err
	}

	cs.Cache.Set(currencyCode, rate)

	return rate, err
}
