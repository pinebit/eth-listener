package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pinebit/eth-listener/token"
)

type Direction int

const (
	Sent Direction = iota
	Received
)

type Transfer struct {
	Direction Direction
	From      common.Address
	To        common.Address
	Value     big.Int
	Token     *token.Token
}
