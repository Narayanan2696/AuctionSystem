package model

import (
	"auction_system/views"
	"fmt"
	"log"
)

func RegisterAuction(auctionId string, auction views.ConductAuction) {
	id := findCount()
	fmt.Println(id)
	insertQ, err := connect.Query("INSERT INTO live_auction(id, auction_id, auctioneer_id, url, bid_value) VALUES(?,?,?,?,?)", id, auctionId, auction.Conductor.Id, auction.Url, auction.BidValue)
	defer insertQ.Close()
	if err != nil {
		log.Println(err.Error)
	}
}

func GetLiveAuction(auctionId string) views.LiveAuction {
	rows, _ := connect.Query("SELECT * FROM live_auction where auction_id = ?", auctionId)
	var live views.LiveAuction
	for rows.Next() {
		rows.Scan(&live.Id, &live.AuctionId, &live.AuctioneerId, &live.Url)
	}
	return live
}

func findCount() int {
	rows, _ := connect.Query("SELECT * FROM live_auction")
	count := 0
	for rows.Next() {
		count += 1
	}
	defer rows.Close()
	return count + 1
}
