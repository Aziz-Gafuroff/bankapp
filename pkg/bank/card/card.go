package card

import (
	"bank/pkg/bank/types"
)

func IssueCard(currency types.Currency, color string, name string) types.Card {
	cardData := types.Card {
	ID: 1000,
	PAN: "5058 xxxx xxxx 0001",
	Balance: 0,
	Currency: currency,
	Color: color,
	Name: name,
	Active: true,

	}

	return cardData
}


const withdrawLimit = 20_000_00

func Withdraw(card types.Card, amount types.Money) types.Card {
	if amount < 0 {
		return card
	}

	if amount > withdrawLimit {
		return card
	}

	if !card.Active {
		return card
	}

	if card.Balance < amount {
		return card
	}

	card.Balance -= amount

	return card
}

const depositLimit = 50_000_00

func Deposit(card *types.Card, amount types.Money) {
	
	if !card.Active {
		return
	}

	if amount > depositLimit {
		return
	}

	card.Balance += amount

}

const maxBonus = 5_000_00

func AddBonus(card *types.Card, percent int, daysInMonth int, daysInYear int) {

	if !card.Active {
		return
	}

	if card.MinBalance < 0 {
		return
	}
	
	Bonus := (int(card.MinBalance) * percent * daysInMonth / daysInYear) / 100

	if Bonus > maxBonus {
		Bonus = maxBonus
	}
	

	card.MinBalance += types.Money(Bonus)
}


func Total(cards []types.Card) types.Money {
	Total := types.Money(0)
	for _, card := range cards {
		Total += card.Balance
	}

	return Total
}

func PaymentSource(cards []types.Card) []types.PaymentSource {
	var source = make ([]types.PaymentSource, 0, len(cards))
	for _, card := range cards {
		
		if card.Balance < 0 {
			continue
		}

		if !card.Active {
			continue
		}

		source = append(source, types.PaymentSource{
			Type: "card",
			Number: string(card.PAN),
			Balance: card.Balance,
		})

	}
	
	return source

}