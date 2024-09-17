package tools

import (
	log "github.com/sirupsen/logrus"
)

type loginDetails struct {
	AuthToken string
	username string
}

type CoinDetails struct {
	Coins int64
	Username string
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *loginDetails
	GetUserCoins(username string) *CoinDetails
	setupDatabse() error
}

func NewDatabase() (*DatabaseInterface, error){
	var database DatabaseInterface = &mockDB{}
	var err error = database.setupDatabse()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &database, nil
}