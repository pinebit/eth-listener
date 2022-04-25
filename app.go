package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type App struct {
	config        *Config
	tokensDB      TokensDB
	accounts      Accounts
	telegram      Telegram
	client        *ethclient.Client
	tokensManager TokensManager
}

func NewApp(config *Config, tokensDB TokensDB, accounts Accounts, telegram Telegram, client *ethclient.Client, tokensManager TokensManager) *App {
	return &App{
		config:        config,
		tokensDB:      tokensDB,
		accounts:      accounts,
		telegram:      telegram,
		client:        client,
		tokensManager: tokensManager,
	}
}
