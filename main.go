package main

import (
	"fmt"
	"log"
	"os"

	"github.com/flipkart/machine-coding-practice/utils"
)

func main() {
	fs := NewFlipStockInstance()

	// add user
	user := fs.NewUser()
	user.Id = 1
	user.Name = "Mayank"
	isSuccess, err := fs.AddUser(user)
	if err != nil {
		log.Println("err: ", err.Error())
		os.Exit(1)
	}
	if !isSuccess {
		log.Println("unable to add user")
		os.Exit(1)
	}
	fmt.Printf("user with id: %v added successfully\n", user.Id)

	// add stock
	stock := fs.NewStock()
	stock.Id = 1
	stock.Name = "SAIL"
	stock.CurrentPrice = 10
	isSuccess, err = fs.AddStock(stock)
	if err != nil {
		log.Println("err: ", err.Error())
		os.Exit(1)
	}
	if !isSuccess {
		log.Println("unable to add stock")
		os.Exit(1)
	}
	fmt.Printf("stock with id: %v added successfully\n", stock.Id)

	// buy stock
	fs.BuyStock(1, 1, 2)

	// sell stock
	fs.SellStock(1, 1, 1)

	// update price
	fs.UpdateStockPrice(0, 1, 20, utils.SYSTEM)

	// add admin user
	user = fs.NewUser()
	user.Id = 2
	user.Name = "Admin"
	isSuccess, err = fs.AddUser(user)
	if err != nil {
		log.Println("err: ", err.Error())
		os.Exit(1)
	}
	if !isSuccess {
		log.Println("unable to add user")
		os.Exit(1)
	}
	fmt.Printf("user with id: %v added successfully\n", user.Id)
	fs.UpdateStockPrice(2, 1, 5, utils.MANUAL)

}
