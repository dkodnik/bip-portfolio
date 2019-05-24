package models

import (
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
	var coins []Coin

	for coin, balance := range s.coins {
		coins = append(coins, Coin{
			Symbol: coin,
			Value:  balance,
		})
	}

	sort.Slice(coins, func(i, j int) bool {
		return coins[i].Value.Cmp(coins[j].Value) == 1
	})

	return coins
}

type Coin struct {
	Symbol string
	Value *big.Int
}