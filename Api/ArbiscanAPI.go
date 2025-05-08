package Api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"tgBot/model"
)

func GetWalletBalance(address string) (float64, error) {
	apiKey := os.Getenv("YOUR_ARBISCAN_API_KEY") // Замените на свой API-ключ
	url := fmt.Sprintf("https://api.arbiscan.io/api?module=account&action=balance&address=%s&tag=latest&apikey=%s", address, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result model.ArbiscanResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	balanceWei, err := strconv.ParseFloat(result.Result, 64)
	if err != nil {
		return 0, err
	}

	// Конвертируем из Wei в ETH
	return balanceWei / 1e18, nil
}
