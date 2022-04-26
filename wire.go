//go:build wireinject

package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/wire"
	"github.com/pinebit/eth-listener/token"
)

func newEthClient(config *Config) (*ethclient.Client, error) {
	return ethclient.Dial(config.EthUrl)
}

func newTokensDB() token.TokensDB {
	return token.NewTokensDB(TokensDBPath)
}

func WireApp(configPath string) (*App, error) {
	wire.Build(NewApp, LoadConfig, NewAccounts, NewTelegram, newEthClient, newTokensDB, token.NewTokensManager)
	return nil, nil
}
