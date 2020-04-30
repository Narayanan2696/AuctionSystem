package model

import (
	"auction_system/lib/errors"
	"auction_system/views"
	"log"
)

func GetAllBidders() ([]views.Bidder, error) {
	rows, err := connect.Query("SELECT * FROM BIDDER")
	var bidders []views.Bidder
	if err != nil {
		log.Println(err.Error)
		return bidders, errors.New(errors.CustomError{404, "Not Found", "bidders not found"})
	}
	for rows.Next() {
		bid := views.Bidder{}
		rows.Scan(&bid.Id, &bid.Name)
		bidders = append(bidders, views.Bidder{bid.Id, bid.Name})
	}
	defer rows.Close()
	return bidders, err
}
