package model

import (
	"auction_system/views"
	"fmt"
	"log"
)

func RegisterBidder(biddingDetails views.ActiveBidders) {
	rows, _ := connect.Query("SELECT * FROM auction_register")
	count := 0
	for rows.Next() {
		count += 1
	}
	fmt.Println(count + 1)
	id := count + 1
	defer rows.Close()
	insertQ, err := connect.Query("INSERT INTO AUCTION_REGISTER VALUES(?,?,?,?)", id, biddingDetails.AuctionId, biddingDetails.BidderId, biddingDetails.Status)
	defer insertQ.Close()
	if err != nil {
		log.Println(err.Error)
	}
}

func ListBidders(auctionId string) ([]views.ActiveBidders, error) {
	rows, err := connect.Query("SELECT * FROM auction_register where auction_id=?", auctionId)
	var activeList []views.ActiveBidders
	for rows.Next() {
		var activeBidders bidderORM
		rows.Scan(&activeBidders.Id, &activeBidders.AuctionId, &activeBidders.BidderId, &activeBidders.Status)
		// fmt.Println(activeBidders)
		activeList = append(activeList, views.ActiveBidders{activeBidders.AuctionId, activeBidders.BidderId, activeBidders.Status})
	}
	if err != nil {
		log.Println(err.Error)
		return []views.ActiveBidders{}, err
	}
	return activeList, err
}

type bidderORM struct {
	Id        int
	AuctionId string
	BidderId  int
	Status    string
}
