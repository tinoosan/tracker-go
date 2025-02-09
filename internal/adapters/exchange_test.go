package adapters_test

import(
	"testing"
	"trackergo/internal/adapters"
)

func TestNewExchangeRateAPI(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		baseURL string
		apiKey  string
		want    *adapters.ExchangeRateAPI
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := adapters.NewExchangeRateAPI(tt.baseURL, tt.apiKey)
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("NewExchangeRateAPI() = %v, want %v", got, tt.want)
			}
		})
	}
}

