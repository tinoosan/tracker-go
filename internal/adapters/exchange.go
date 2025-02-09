package adapters

import (
	"encoding/json"
	"fmt"
	"net/http"
	vo "trackergo/internal/domain/valueobjects"

)

type ExchangeRateAPI struct {
	BaseURL string
	APIKey  string
}

func NewExchangeRateAPI(baseURL, apiKey string) *ExchangeRateAPI {
	return &ExchangeRateAPI{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (api *ExchangeRateAPI) GetExchangeRate(baseCurrency, targetCurrency string) (*vo.Ratio, error) {

	url := fmt.Sprintf("%s/latest?base=%s&symbols=%s&apikey=%s", api.BaseURL, baseCurrency, targetCurrency, api.APIKey)

  resp, err := http.Get(url)
  if err != nil {
    return &vo.Ratio{}, fmt.Errorf("failed to fetch exchange rate: %w", err)
  }

  var data struct {
    Rates map[string]float64 `json:"rates"`
  }

  if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
    return &vo.Ratio{}, fmt.Errorf("failed to decode exchange rate response: %w", err)
  }

  rate, ok := data.Rates[targetCurrency]
  if !ok {
    return &vo.Ratio{}, fmt.Errorf("exchange rate for %s not found", targetCurrency)
  }

  ratio, err := vo.NewRatio(rate)
  if err != nil {
    return &vo.Ratio{}, err
  }

  return ratio, nil

}
