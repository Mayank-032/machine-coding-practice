package user

import (
	"errors"

	"github.com/flipkart/machine-coding-practice/stock"
	"github.com/flipkart/machine-coding-practice/utils"
)

type User struct {
	Id     int                              `json:"id"`
	Name   string                           `json:"name"`
	Role   string                           `json:"role"`
	Stocks map[int]userBasedStockProperties `json:"stocks"`
}

type userBasedStockProperties struct {
	Quantity     int     `json:"quantity"`
	TotalPrice   float64 `json:"total_price"`
	AveragePrice float64 `json:"average_price"`
}

func New() *User {
	return &User{}
}

func (u *User) Buy(stock stock.Stock, quantity int) error {
	if u.Role != utils.USER {
		return errors.New("user not allowed to perform this operation")
	}

	var existingStockProperties userBasedStockProperties
	var ok bool
	if existingStockProperties, ok = u.Stocks[stock.Id]; !ok {
		existingStockProperties = userBasedStockProperties{}
	}

	existingStockProperties.TotalPrice += stock.CurrentPrice
	existingStockProperties.Quantity += quantity
	existingStockProperties.AveragePrice = existingStockProperties.TotalPrice / float64(existingStockProperties.Quantity)

	u.Stocks[stock.Id] = existingStockProperties
	return nil
}

func (u *User) Sell(stock stock.Stock, quantity int) error {
	if u.Role != utils.USER {
		return errors.New("user not allowed to perform this operation")
	}

	var existingStockProperties userBasedStockProperties
	var ok bool
	if existingStockProperties, ok = u.Stocks[stock.Id]; !ok {
		return errors.New("invalid operation. stock does not exists with user")
	}

	if existingStockProperties.Quantity < quantity {
		return errors.New("invalid operation. current stock quantity is less than selected sell quantity")
	}

	existingStockProperties.TotalPrice -= stock.CurrentPrice
	existingStockProperties.Quantity -= quantity
	existingStockProperties.AveragePrice = existingStockProperties.TotalPrice / float64(existingStockProperties.Quantity)

	u.Stocks[stock.Id] = existingStockProperties
	return nil
}
