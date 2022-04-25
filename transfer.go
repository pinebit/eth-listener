package main

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
	Token     *Token
}
