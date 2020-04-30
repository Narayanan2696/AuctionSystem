package model

import (
	"auction_system/views"
	"log"
)

func InsertWinningBid(bid views.Bid, auctionId string) {
	insertQ, err := connect.Query("INSERT INTO WINNING_BID VALUES(?,?,?,?)", findId(), bid.BidderId, auctionId, bid.BiddingAmount)
	defer insertQ.Close()
	if err != nil {
		log.Println(err.Error)
	}
}

func GetWinningBid(auctionId string) (views.Bid, error) {
	rows, err := connect.Query("SELECT * FROM WINNING_BID where auction_id=?", auctionId)
	var bid views.WinningBid
	for rows.Next() {
		rows.Scan(&bid.Id, &bid.BidderId, &bid.AuctionId, &bid.BiddingAmount)
	}
	defer rows.Close()
	if err != nil {
		log.Println(err.Error)
		return views.Bid{}, err
	}
	return views.Bid{bid.BidderId, bid.BiddingAmount}, err
}

func findId() int {
	rows, _ := connect.Query("SELECT * FROM WINNING_BID")
	count := 0
	for rows.Next() {
		count += 1
	}
	defer rows.Close()
	return count + 1
}
