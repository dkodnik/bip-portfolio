package helpers

import (
	"github.com/danil-lashin/bip-portfolio/config"
	"github.com/dustin/go-humanize"
	"math/big"
)

var pipInBip = big.NewFloat(1000000000000000000)

func StrToBig(balance string) *big.Int {
	bigInt, success := big.NewInt(0).SetString(balance, 10)
	if success != true {
		panic("Failed to decode " + balance)
	}

	return bigInt
}

func PipToBip(value *big.Int) string {
	floatValue := new(big.Float).SetPrec(500).SetInt(value)
	f, _ := new(big.Float).SetPrec(500).Quo(floatValue, pipInBip).Float64()

	return humanize.FormatFloat("# ###.##", f)
}

func BipToUsd(value *big.Int) string {
	floatValue := new(big.Float).SetPrec(500).SetInt(value)
	f, _ := new(big.Float).SetPrec(500).Quo(floatValue, pipInBip).Float64()

	return "$" + humanize.FormatFloat("# ###.##", f*config.Get().BipPriceUSD)
}
