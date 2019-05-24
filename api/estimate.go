package api

type EstimateResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  struct {
		WillGet    string `json:"will_get"`
		Commission string `json:"commission"`
	} `json:"result"`
}
