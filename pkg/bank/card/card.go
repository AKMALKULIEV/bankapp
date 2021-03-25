package card

import (
	"bank/pkg/bank/types"
	// "fmt"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	card := types.Card{
		Id:       1000,
		Pan:      "5058 xxxx xxxx 0001",
		Balance:  0,
		Currency: currency,
		Color:    color,
		Name:     name,
		Active:   true,
	}
	return card
}
const withdrawlimit = types.Money(20_000_00)
func Withdraw(card *types.Card, amount types.Money){
	if amount < 0 {
	 return
	}
	if amount > withdrawlimit{
		return
	}
	if !(*card).Active{
		return
	}
	if (*card).Balance < amount{
		return
	}
	(*card).Balance -= amount	
}

const depositLimit = 50_000_00
 func Deposit(card *types.Card, amount types.Money) {
	 if !(*card).Active{
		 return 
	 }
	 if amount> depositLimit{
		 return
	 }
	 if amount < 0{
		 return
	 }
	 (*card).Balance += amount
 }
//  const bonusLimit = 5000
 
 func AddBonus(card *types.Card, percent int, daysInMonth int, daysiInYear int)  {
	if card.Balance <=0{
		return
	}
	if !(*card).Active{
		 return
	 }
	 if (*card).MinBalance < 0{
		 return
	 }
	 if (*card).MinBalance == 0{
		return
	}

	  (*card).MinBalance = card.MinBalance*types.Money(percent)
	  (*card).MinBalance = card.MinBalance/100
	  (*card).MinBalance = card.MinBalance*types.Money(daysInMonth)
	  (*card).MinBalance = card.MinBalance/types.Money(daysiInYear)
	  if (*card).MinBalance>5_000_00{
        // (*card).MinBalance = card.MinBalance-(card.MinBalance-5_000)
		(*card).MinBalance = 0
	  }
	  (*card).Balance += card.MinBalance

 }
func Total(cards []types.Card)  types.Money{
	
	sum := types.Money(0)
	for _, card := range cards {
		if card.Balance <= 0 {
		continue
		}
		if !card.Active {
			continue
			}
		sum += card.Balance	
	}
	
	return sum
}