package api

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

type Node struct {
	nodeUrl string
}

func NewNodeAPI(nodeUrl string) *Node {
	return &Node{
		nodeUrl: nodeUrl,
	}
}

func (n *Node) GetCandidates() *CandidatesResponse {
	bytes := n.httpGet("/candidates?include_stakes=true")

	candidates := &CandidatesResponse{}
	if err := json.Unmarshal(bytes, candidates); err != nil {
		panic(err)
	}

	return candidates
}

func (n *Node) GetAddress(address string) *AddressResponse {
	bytes := n.httpGet("/address?address=" + address)

	response := &AddressResponse{}
	if err := json.Unmarshal(bytes, response); err != nil {
		panic(err)
	}

	return response
}

func (n *Node) EstimateCoinSell(coinToSell string, coinToBuy string, valueToSell string) *EstimateResponse {
	bytes := n.httpGet("/estimate_coin_sell?coin_to_sell=" + coinToSell + "&coin_to_buy=" + coinToBuy + "&value_to_sell=" + valueToSell)

	estimate := &EstimateResponse{}
	if err := json.Unmarshal(bytes, estimate); err != nil {
		panic(err)
	}

	return estimate
}

func (n *Node) httpGet(endpoint string) []byte {
	_, bytes, err := fasthttp.Get(nil, n.nodeUrl+endpoint)
	if err != nil {
		panic(err)
	}

	return bytes
}
