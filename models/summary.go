package models

import (
	"github.com/danil-lashin/bip-portfolio/api"
	"github.com/danil-lashin/bip-portfolio/config"
	"github.com/danil-lashin/bip-portfolio/helpers"
	"math/big"
	"sort"
)

type Summary struct {
	coins map[string]*big.Int
}

func NewSummary() *Summary {
	return &Summary{
		coins: map[string]*big.Int{},
	}
}

func (s *Summary) AddCoin(coin string, value *big.Int) {
	if s.coins[coin] == nil {
		s.coins[coin] = big.NewInt(0)
	}

	s.coins[coin].Add(s.coins[coin], value)
}

func (s *Summary) GetCoins() []Coin {
	nodeApi := api.NewNodeAPI(config.Get().NodeURL)

	var coins []Coin
	for coin, balance := range s.coins {
		bipValue := balance
		if coin != "BIP" {
			estimate := nodeApi.EstimateCoinSell(coin, "BIP", balance.String())
			bipValue = helpers.StrToBig(estimate.Result.WillGet)
		}

		coins = append(coins, Coin{
			Symbol:   coin,
			Value:    balance,
			BipValue: bipValue,
		})
	}

	sort.Slice(coins, func(i, j int) bool {
		return coins[i].BipValue.Cmp(coins[j].BipValue) == 1
	})

	return coins
}

type Coin struct {
	Symbol   string
	Value    *big.Int
	BipValue *big.Int
}
