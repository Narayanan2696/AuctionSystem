package controller

import (
	"auction_system/lib"
	"auction_system/lib/render"
	"auction_system/service"
	"auction_system/views"
	"encoding/json"
	"net/http"
)

func PostAuctionRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			auctionData := views.AuctionRequest{}
			json.NewDecoder(r.Body).Decode(&auctionData)

			auctioneer, err := lib.FetchAuctioneer(auctionData.AuctionId)
			// bidders, err := lib.FetchBidders()
			service.MaxBidder(auctioneer, auctionData.AuctionId)
			render.JSON(w, err, auctioneer)
		}
	}
}
