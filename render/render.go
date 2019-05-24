package render

import (
	"fmt"
	"github.com/danil-lashin/bip-portfolio/config"
	"github.com/danil-lashin/bip-portfolio/helpers"
	"github.com/danil-lashin/bip-portfolio/models"
	"github.com/olekukonko/tablewriter"
	"math/big"
	"os"
)

func Render(s *models.Summary) {
	cfg := config.Get()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.SetColumnAlignment([]int{tablewriter.ALIGN_LEFT, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_RIGHT, tablewriter.ALIGN_RIGHT})
	table.SetHeader([]string{"Coin", "Value", "Bip Value", "USD Value (1 bip = $" + fmt.Sprintf("%.2f", cfg.BipPriceUSD) + ")"})
	totalOwned := big.NewInt(0)

	var data [][]string
	for _, coin := range s.GetCoins() {
		data = append(data, []string{coin.Symbol, helpers.PipToBip(coin.Value), helpers.PipToBip(coin.BipValue), helpers.BipToUsd(coin.BipValue)})
		totalOwned.Add(totalOwned, coin.BipValue)
	}

	table.AppendBulk(data)
	table.SetFooter([]string{"", "Total", helpers.PipToBip(totalOwned), helpers.BipToUsd(totalOwned)})
	table.Render()
}
