package service

import (
	"auction_system/views"
	"fmt"
)

func MaxBidder(auctioner views.ConductAuction, auctionId string) {
	// bidders, err := lib.FetchBidders()
	// var bidderDetails []views.BidderResponse
	// if err != nil {
	// 	log.Println(err.Error())
	// }
	// biddersChannel := make(chan []views.BidderResponse)
	// delayIncrementor := 0
	// for len(bidderDetails) > 0 {
	// 	delay_timer := time.NewTicker(time.Duration(auctioner.DelayMS) * time.Millisecond)
	// 	go getBidderDetails(biddersChannel, delay_timer)

	// 	delayIncrementor +=
	// }

	fmt.Println("max bidder auction id:", auctionId)
}
