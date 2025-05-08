package model

// ArbiscanResponse представляет структуру ответа от Arbiscan API
type ArbiscanResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  string `json:"result"`
}
