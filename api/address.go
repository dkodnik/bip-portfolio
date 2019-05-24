package api

type AddressResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  struct {
		Balance          map[string]string `json:"balance"`
		TransactionCount string            `json:"transaction_count"`
	} `json:"result"`
}
