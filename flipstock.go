package main

import (
	"errors"
	"sync"

	"github.com/flipkart/machine-coding-practice/stock"
	"github.com/flipkart/machine-coding-practice/user"
	"github.com/flipkart/machine-coding-practice/utils"
)

type FlipStock struct {
	Users  map[int]user.User   `json:"users"`
	Stocks map[int]stock.Stock `json:"stocks"`
}

func NewFlipStockInstance() *FlipStock {
	var f *FlipStock

	sync.OnceFunc(func() {
		f = &FlipStock{
			Users:  make(map[int]user.User, 0),
			Stocks: make(map[int]stock.Stock, 0),
		}
	})
	return f
}

func (fs *FlipStock) NewUser() *user.User {
	return user.New()
}

func (fs *FlipStock) NewStock() *stock.Stock {
	return stock.New()
}

func (fs *FlipStock) AddUser(user *user.User) (bool, error) {
	if err := user.Validate(); err != nil {
		return false, err
	}

	fs.Users[user.Id] = *user
	return true, nil
}

func (fs *FlipStock) AddStock(stock *stock.Stock) (bool, error) {
	if err := stock.Validate(); err != nil {
		return false, err
	}

	fs.Stocks[stock.Id] = *stock
	return true, nil
}

func (fs *FlipStock) BuyStock(userId, stockId, quantity int) error {
	if _, ok := fs.Users[userId]; ok {
		return errors.New("user does not exists")
	}

	if _, ok := fs.Stocks[stockId]; ok {
		return errors.New("stock does not exists")
	}

	user := fs.Users[userId]
	err := user.Buy(fs.Stocks[stockId], quantity)
	if err != nil {
		return err
	}
	fs.Users[userId] = user

	return nil
}

func (fs *FlipStock) SellStock(userId, stockId, quantity int) error {
	if _, ok := fs.Users[userId]; ok {
		return errors.New("user does not exists")
	}

	if _, ok := fs.Stocks[stockId]; ok {
		return errors.New("stock does not exists")
	}

	user := fs.Users[userId]
	err := user.Sell(fs.Stocks[stockId], quantity)
	if err != nil {
		return err
	}
	fs.Users[userId] = user

	return nil
}

func (fs *FlipStock) UpdateStockPrice(userId, stockId int, newPrice float64, operationType string) error {
	if _, ok := fs.Stocks[stockId]; ok {
		return errors.New("stock does not exists")
	}

	stock := fs.Stocks[stockId]

	if operationType == utils.MANUAL {
		if _, ok := fs.Users[userId]; ok {
			return errors.New("user does not exists")
		}

		user := fs.Users[userId]
		if user.Role != utils.ADMIN {
			return errors.New("invalid operation. user must be an admin")
		}
	}

	stock.UpdateCurrentPrice(newPrice)
	fs.Stocks[stockId] = stock

	return nil
}
