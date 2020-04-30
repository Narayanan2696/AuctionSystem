package controller

import (
	"github.com/gorilla/mux"
)

func Register() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/auction", PostAuctionRequest())
	router.HandleFunc("/auctioneer", GetConductAuctioneer())
	router.HandleFunc("/bidders", GetBidders())
	router.HandleFunc("/list_bidders/{auction_id}", ListOnlineBidders())
	router.HandleFunc("/register_bidders/{auction_id}", RegisterBidders())
	router.HandleFunc("/bid/{auction_id}", WinningBid())
	return router
}
