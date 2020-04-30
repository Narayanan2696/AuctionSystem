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
	return router
}
