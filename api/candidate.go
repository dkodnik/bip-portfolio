package api

type CandidatesResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Id      string `json:"id"`
	Result  []struct {
		RewardAddress string `json:"reward_address"`
		OwnerAddress  string `json:"owner_address"`
		TotalStake    string `json:"total_stake"`
		PubKey        string `json:"pub_key"`
		Commission    string `json:"commission"`
		Stakes        []struct {
			Owner    string `json:"owner"`
			Coin     string `json:"coin"`
			Value    string `json:"value"`
			BipValue string `json:"bip_value"`
		} `json:"stakes"`
		Status         int    `json:"status"`
		CreatedAtBlock string `json:"created_at_block"`
	} `json:"result"`
}
