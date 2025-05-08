package Api

import (
	"encoding/json"
	"net/http"
	"tgBot/model"
)

func GetETHPrice() (float64, error) {
	url := "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result model.CoinGeckoResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	return result.Ethereum.USD, nil
}
