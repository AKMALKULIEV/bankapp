package payment

import (
	"bank/pkg/bank/types"
	"fmt"
)

func ExampleMax() {
	payments := []types.Payment{
		{
			ID:     1,
			Amount: 500000,
		},
		{
			ID:     2,
			Amount: 500001,
		},
		{
			ID:     3,
			Amount: 500001,
		},
	}
	max := Max(payments)
	fmt.Println(max.ID)
	//Output:
	//2
}
