package payment

import (
	"bank/pkg/bank/types"
	
	// "fmt"
)

func Max(payments []types.Payment)  types.Payment{
    max := types.Payment{}
	for _, payment := range payments {
		if max.Amount < payment.Amount{
			max = payment
			payments[0] = max
		}

		
	}
    return max 	
}
