package main

import (
	"fmt"

	"github.com/BehroozPayoon/crypto-exchange/orderbook"
)

func main() {
	fmt.Println("It's a crypto exchange system!!!")
}

type Market string

const (
	MarketEth Market = "ETH"
)

type Exchange struct {
	orderbooks map[Market]*orderbook.Orderbook
}

func NewExchange() *Exchange {
	return &Exchange{
		orderbooks: make(map[Market]*orderbook.Orderbook),
	}
}
