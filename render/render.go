package render

import (
	"fmt"
	"github.com/danil-lashin/bip-portfolio/api"
	"github.com/danil-lashin/bip-portfolio/config"
	"github.com/danil-lashin/bip-portfolio/helpers"
	"github.com/danil-lashin/bip-portfolio/models"
	"github.com/olekukonko/tablewriter"
	"math/big"
	"os"
)

func Render(s *models.Summary) {
	cfg := config.Get()
	nodeApi := api.NewNodeAPI(cfg.NodeURL)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_RIGHT})
	table.SetHeader([]string{"Coin", "Value", "Bip Value", "USD Value (1 bip = $" + fmt.Sprintf("%.2f", cfg.BipPriceUSD) + ")"})
	totalOwned := big.NewInt(0)

	var data [][]string
	for _, coin := range s.GetCoins() {
		if coin.Symbol == "BIP" {
			data = append(data, []string{coin.Symbol, helpers.PipToBip(coin.Value), helpers.PipToBip(coin.Value), helpers.BipToUsd(coin.Value)})
			totalOwned.Add(totalOwned, coin.Value)
			continue
		}

		estimate := nodeApi.EstimateCoinSell(coin.Symbol, "BIP", coin.Value.String())

		bipValue := helpers.StrToBig(estimate.Result.WillGet)
		data = append(data, []string{coin.Symbol, helpers.PipToBip(coin.Value), helpers.PipToBip(bipValue), helpers.BipToUsd(bipValue)})
		totalOwned.Add(totalOwned, bipValue)
	}

	table.AppendBulk(data)
	table.SetFooter([]string{"", "Total", helpers.PipToBip(totalOwned), helpers.BipToUsd(totalOwned)})
	table.Render()
}
