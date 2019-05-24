package main

import (
	"github.com/danil-lashin/bip-portfolio/api"
	"github.com/danil-lashin/bip-portfolio/config"
	"github.com/danil-lashin/bip-portfolio/helpers"
	"github.com/danil-lashin/bip-portfolio/models"
	"github.com/danil-lashin/bip-portfolio/render"
)

func main() {
	cfg := config.Get()
	nodeApi := api.NewNodeAPI(cfg.NodeURL)
	summary := models.NewSummary()

	candidates := nodeApi.GetCandidates()

	for _, address := range cfg.Addresses {
		response := nodeApi.GetAddress(address)

		for coin, balance := range response.Result.Balance {
			summary.AddCoin(coin, helpers.StrToBig(balance))
		}

		for _, candidate := range candidates.Result {
			for _, stake := range candidate.Stakes {
				if stake.Owner == address {
					summary.AddCoin(stake.Coin, helpers.StrToBig(stake.Value))
				}
			}
		}
	}

	render.Render(summary)
}
