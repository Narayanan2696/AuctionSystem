package model

import (
	"auction_system/lib/errors"
	"auction_system/views"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GetAuctioneer() (views.Auctioneer, error) {
	rows, _ := connect.Query("SELECT * FROM AUCTIONEER")
	count := 0
	auctioneer := views.Auctioneer{}
	for rows.Next() {
		count += 1
	}
	defer rows.Close()
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(count-1) + 1
	fmt.Println("id:", id)
	auctioner, err := connect.Query("SELECT * FROM AUCTIONEER where id = ?", id)
	if err != nil {
		log.Println(err.Error)

		return auctioneer, errors.New(errors.CustomError{404, "Not Found", "auctioner id:" + strconv.Itoa(id) + "not found"})
	}
	for auctioner.Next() {
		auctioner.Scan(&auctioneer.Id, &auctioneer.Name)
	}
	defer auctioner.Close()

	return auctioneer, err
}
