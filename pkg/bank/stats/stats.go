package stats

import (
	"github.com/kodirovrukhillo/bank-avg/pkg/bank/types"
)

func Avg(payments []types.Payment) types.Money {
	var average types.Money
	for _, payment := range payments {
		if payment.Amount > 0 {
			average += payment.Amount

		}

	}
	average1 := average / types.Money(len(payments))
	return average1
}
