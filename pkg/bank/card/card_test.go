package card

import (
	"bank/pkg/bank/types"
	"fmt"
)


func ExampleWithdraw_positive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 10_000_00)

	fmt.Println(result.Balance)

	//Output: 1000000
}


func ExampleWithdraw_noMoney() {
	result := Withdraw(types.Card{Balance: 0, Active: true}, 10_000_00)

	fmt.Println(result.Balance)

	//Output: 0
}

func ExampleWithdraw_inActive() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: false}, 10_000_00)

	fmt.Println(result.Balance)

	//Output: 2000000
}

func ExampleWithdraw_limit() {
	result := Withdraw(types.Card{Balance: 20_000_00, Active: true}, 21_000_00)

	fmt.Println(result.Balance)

	//Output: 2000000
}

func ExampleDeposit_positive() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 10_000_00)
	
	fmt.Println(card.Balance)

	//Output: 3000000

}

func ExampleDeposit_inActive() {
	card := types.Card{Balance: 20_000_00, Active: false}
	Deposit(&card, 10_000_00)
	
	fmt.Println(card.Balance)

	//Output: 2000000
}

func ExampleDeposit_limit() {
	card := types.Card{Balance: 20_000_00, Active: true}
	Deposit(&card, 55_000_00)
	
	fmt.Println(card.Balance)

	//Output: 2000000

}

func ExampleAddBonus_positive() {
	card := types.Card{MinBalance: 10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.MinBalance)

	// Output: 1002465

}

func ExampleAddBonus_inActive() {
	card := types.Card{MinBalance: 10_000_00, Active: false}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.MinBalance)

	// Output: 1000000

}

func ExampleAddBonus_negativeBalance() {
	card := types.Card{MinBalance: -10_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.MinBalance)

	// Output: -1000000

}

func ExampleAddBonus_moreMaxBonus() {
	card := types.Card{MinBalance: 21_000_000_00, Active: true}
	AddBonus(&card, 3, 30, 365)

	fmt.Println(card.MinBalance)

	// Output: 2100500000

}


func ExampleTotal() {
	cards := []types.Card{
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
		{
			Balance: 10_000_00,
			Active: true,
		},
	}

	fmt.Println(Total(cards))

	// Output: 3000000
}

func ExamplePaymentSource() {
	cards := []types.Card{
		{
			Balance: 20_000_00,
			PAN: "5058 xxxx xxxx 8888",
			Active: true,
		},
		{
			Balance: 10_000_00,
			PAN: "5058 xxxx xxxx 4444",
			Active: false,
		},
		{
			Balance: 50_000_00,
			PAN: "5058 xxxx xxxx 9999",
			Active: true,
		},
	}

	fmt.Println(PaymentSource(cards))

	//Output: [{card 5058 xxxx xxxx 8888 2000000} {card 5058 xxxx xxxx 9999 5000000}]
}