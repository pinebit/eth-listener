//go:build wireinject

package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/wire"
)

func newEthClient(config *Config) (*ethclient.Client, error) {
	return ethclient.Dial(config.EthUrl)
}

func WireApp(configPath string) (*App, error) {
	wire.Build(NewApp, LoadConfig, NewTokensDB, NewAccounts, NewTelegram, newEthClient, NewTokensManager)
	return nil, nil
}
