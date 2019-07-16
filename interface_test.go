package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPayByWallet(t *testing.T) {
	amount := 50.0
	payment := 30.0
	w := &Wallet{amount}
	w.Pay(float32(payment))

	assert.Equal(t, w.funds, amount-payment)
}

func TestPayByWalletGivesError(t *testing.T) {
	amount := 50.0
	payment := 60.0
	w := &Wallet{amount}
	err := w.Pay(float32(payment))
	assert.Error(t, err)
}

func TestBuyByWallet(t *testing.T) {
	amount := 50
	payment := 60
	w := &Wallet{amount}
	_ = Buy(w, payment)
	assert.Equal(t, w.funds, amount-payment)
}

func TestBuyByCreditCard(t *testing.T) {
	amount := 20
	payment := 30
	c := &CreditCard{funds: amount}
	_ = Buy(c, payment)
	assert.Equal(t, c.funds, amount-payment)
}
func TestBuyByBitcoin(t *testing.T) {
	amount := 23.3
	payment := 67
	b := &Bitcoin{funds: amount, owner: "Karina"}
	_ = Buy(b, payment)
	fmt.Println(b)
}

func TestCheckAndBuyWallet(t *testing.T) {
	amount := 50
	w := &Wallet{amount}
	checkAndBuy(w, amount)
	assert.Equal(t, w.funds, amount)
}

func TestCheckPaymentType(t *testing.T) {
	amount := 50
	c := &CreditCard{amount, "KArina", time.Time{}, nil}
	fmt.Println(CheckPaymentType(c))
}

func TestHashTranaction(t *testing.T) {
	test := "hello"
	fmt.Println(HashTransaction(test))
}

func TestBitcoin(t *testing.T) {
	b := &Bitcoin{funds: 20, owner: "KArina"}
	fmt.Println(b.transactions[0].previousTransaction)
}
func TestGetFundsCredit(t *testing.T) {
	w := &Wallet{50}
	fmt.Println(w.GetFunds())
}
