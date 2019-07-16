package main

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"
)

var (
	newTransaction Transaction
)

type Wallet struct {
	funds float32
}
type CreditCard struct {
	funds     float32
	owner      string
	expireTime time.Time
	bonuses    []Bonus
}
type Bitcoin struct {
	funds        float32
	owner        string
	transactions []Transaction
}
type Transaction struct {
	sumOfTransaction    float32
	transactionTime     time.Time
	previousTransaction string
}

func HashTransaction(previousTransactionToString string) string {
	h := sha1.New()
	h.Write([]byte(previousTransactionToString))
	previousTransaction := h.Sum(nil)
	return hex.EncodeToString(previousTransaction)
}

type Bonus struct {
	bonusName   string
	description string
}

type Payer interface {
	Pay(float32) error
}

type Funder interface {
	GetFunds() int
}

type PayFunder interface {
	Payer
	Funder
}

func (w *Wallet) Pay(amount float32) error {
	if w.funds < amount {
		return errors.New("not enough founds")
	}
	w.funds -= amount
	return nil
}

func (c *CreditCard) Pay(amount float32) error {
	if c.funds < amount {
		return errors.New("not enough founds")
	}
	c.funds -= amount
	return nil
}
func (b *Bitcoin) Pay(amount float32) error {
	beforeTransaction := b.funds
	if b.funds < amount {
		return errors.New("not enough founds")
	}
	b.funds -= amount
	if len(b.transactions) != 0 {
		newTransaction = Transaction{
			b.funds - beforeTransaction,
			time.Now(),
			convertToNextTransaction(b.transactions[len(b.transactions)-1].sumOfTransaction, b.transactions[len(b.transactions)-1].transactionTime, b.transactions[len(b.transactions)-1].previousTransaction),
		}

		b.transactions = append(b.transactions, newTransaction)
	} else {
		newTransaction = Transaction{
			b.funds - beforeTransaction,
			time.Now(),
			"",
		}
		b.transactions = append(b.transactions, newTransaction)
	}

	return nil
}

func Buy(p Payer, amount int) error {
	err := p.Pay(float32(amount))
	if err != nil {
		return err
	}
	return nil
}
func checkAndBuy(p PayFunder, amount int) error {
	if p.GetFunds() <= 0 {
		fmt.Println("пополните счет")
	}
	err := p.Pay(float32(amount))
	if err != nil {
		return err
	}
	return nil
}

func (w *Wallet) GetFunds() float32 {
	return w.funds
}

func CheckPaymentType(p Payer) interface{} {
	switch p.(type) {
	case *Wallet:
		fmt.Println("ты пользуешься кошельком")
		return p.(*Wallet).funds
	case *CreditCard:
		fmt.Println("ты пользуешься кредиткой ")
		fmt.Println(p.(*CreditCard).owner)
		return p.(*CreditCard).funds
	case *Bitcoin:
		fmt.Println("ты пользуешься биткойном, хы ")
		fmt.Println(p.(*Bitcoin).owner)
		return p.(*Bitcoin).funds

	default:
		return nil
	}
}

func convertToNextTransaction(sumOfTransaction float32, transactionTime time.Time, previousTransaction string) string {
	return HashTransaction(fmt.Sprintf(strconv.Itoa(int(sumOfTransaction)), transactionTime, previousTransaction))
}
