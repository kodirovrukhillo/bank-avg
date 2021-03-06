package stats

import (
	"fmt"

	"github.com/kodirovrukhillo/bank-avg/pkg/bank/types"
)

func ExampleAvg() {
	payment := []types.Payment{
		{
			ID:       11111,
			Amount:   10_000_00,
			Category: "Restraunt",
		},
		{
			ID:       1234,
			Amount:   18_000_00,
			Category: "Pharmacy",
		},
		{
			ID:       10000,
			Amount:   5_000_00,
			Category: "Pharmacy",
		},
	}
	result := Avg(payment)
	fmt.Println(result)
}

// Output:
// 4300000
