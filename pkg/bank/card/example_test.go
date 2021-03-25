package card

import (
	"bank/pkg/bank/types"
	"fmt"
)

// func ExampleWithdraw_positive() {
// 	card := types.Card{Balance: 2000000, Active: true}
// 	Withdraw(&card, 10_000_00)
// 	fmt.Println(card.Balance)
// 	// Output: 1000000
// }
// func ExampleWithdraw_noMoney() {
// 	card := types.Card{Balance: 0, Active: true}
// 	Withdraw(&card, 50_000_00)
// 	fmt.Println(card.Balance)
// 	// Output: 0
// }
// func ExampleWithdraw_inactive() {
// 	card := types.Card{Balance: 500000, Active: false}
// 	Withdraw(&card, 50_000_00)
// 	fmt.Println(card.Balance)
// 	// Output: 500000
// }
// func ExampleWithdraw_limit() {
// 	card := types.Card{Balance: 1100000, Active: false}
// 	Withdraw(&card, 40_000_00)
// 	fmt.Println(card.Balance)
// 	// Output: 1100000
// }

// func ExampleDeposit_positive() {
// 	card := types.Card{Balance: 5000, Active: true}
// 	Deposit(&card, 5404)
// 	fmt.Println(card.Balance)
// 	// Output:
// 	// 10404

// }
// func ExampleDeposit_inactive() {
// 	card := types.Card{Balance: 5000, Active: false}
// 	Deposit(&card, 5404)
// 	fmt.Println(card.Balance)
// 	// Output:
// 	// 5000

// }
// func ExampleDeposit_limit() {
// 	card := types.Card{Balance: 5000, Active: true}
// 	Deposit(&card, 54040000)
// 	fmt.Println(card.Balance)
// 	// Output:
// 	// 5000

// }
// func ExampleAddBonus_positive() {
// 	card := types.Card{Balance: 100, MinBalance: 10000, Active: true}
// 	AddBonus(&card, 3, 30, 364)
// 	fmt.Println(card.Balance)
// 	// Output:
// 	// 124
// }
// func ExampleAddBonus_inactive() {
// 	card := types.Card{Balance: 100, MinBalance: 10000, Active: false}
// 	AddBonus(&card, 3, 30, 364)
// 	fmt.Println(card.Balance)
// 	// Output:
// 	// 100
// }
// func ExampleAddBonus_lessthan() {
// 	card := types.Card{Balance: -1, MinBalance: 10000, Active: true}
// 	AddBonus(&card, 3, 30, 364)
// 	fmt.Println(card.Balance)
// 	// Output:
// 	// -1
// }
// func ExampleAddBonus_limit() {
// 	card := types.Card{Balance: 0, MinBalance: 10_000_000_00, Active: true}
// 	AddBonus(&card, 3, 30, 364)
// 	fmt.Println(card.Balance)
// 	// Output:
// 	// 0
// }

func ExampleTotal() {
	fmt.Println(Total([]types.Card{
	   {
		Balance: 100000,
		Active:  true,
	    },      
	}))
	fmt.Println(Total([]types.Card{
		{
		Balance: 100000,
		Active:  true,
	    },
		{
		Balance: 200000,
		Active:  true,
	    },       
	}))
	fmt.Println(Total([]types.Card{
		{
		 Balance: 100000,
		 Active:  false,
		 },      
	 }))
	 fmt.Println(Total([]types.Card{
		{
		 Balance: -100000,
		 Active:  true,
		 },      
	 }))

	//Output:
	// 100000
	// 300000
	// 0
	// 0

}
