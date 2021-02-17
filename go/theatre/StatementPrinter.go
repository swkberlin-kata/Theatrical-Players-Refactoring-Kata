package theatre

import (
	"fmt"
	"math"

	"github.com/leekchan/accounting"
)

type StatementPrinter struct{}

func (StatementPrinter) Print(invoice Invoice, plays map[string]Play) (string, error) {
	totalAmount := 0
	volumeCredits := 0
	result := fmt.Sprintf("Statement for %s\n", invoice.Customer)

	for _, perf := range invoice.Performances {
		play := plays[perf.PlayID]

		thisAmount, e := computePrice(play.Type, perf)
		if e != nil {
			return "", e
		}

		// add volume credits
		volumeCredits += int(math.Max(float64(perf.Audience)-30, 0))
		// add extra credit for every ten comedy attendees
		if play.Type == "comedy" {
			volumeCredits += int(math.Floor(float64(perf.Audience) / 5))
		}

		// print line for this order
		result += fmt.Sprintf("  %s: %s (%d seats)\n", play.Name, formatMoney(thisAmount), perf.Audience)
		totalAmount += thisAmount
	}
	result += fmt.Sprintf("Amount owed is %s\n", formatMoney(totalAmount))
	result += fmt.Sprintf("You earned %d credits\n", volumeCredits)
	return result, nil
}

func formatMoney(totalAmount int) string {
	ac2 := accounting.Accounting{Symbol: "$", Precision: 2}
	return ac2.FormatMoney(float64(totalAmount) / 100)
}

func computePrice(playType string, perf Performance) (int, error) {
	thisAmount := 0

	switch playType {
	case "tragedy":
		thisAmount = 40000
		if perf.Audience > 30 {
			thisAmount += 1000 * (perf.Audience - 30)
		}
	case "comedy":
		thisAmount = 30000
		if perf.Audience > 20 {
			thisAmount += 10000 + 500*(perf.Audience-20)
		}
		thisAmount += 300 * perf.Audience
	default:
		return 0, fmt.Errorf("unknown type: %s", playType)
	}
	return thisAmount, nil
}
