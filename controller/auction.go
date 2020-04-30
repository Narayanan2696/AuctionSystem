package controller

import (
	"auction_system/lib/render"
	"auction_system/model"
	"auction_system/service"
	"auction_system/views"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
)

func PostAuctionRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			auctionData := views.AuctionRequest{}
			json.NewDecoder(r.Body).Decode(&auctionData)

			// auctioneer, err := lib.FetchAuctioneer(auctionData.AuctionId)
			delay, _ := strconv.Atoi(os.Getenv("DELAY_MS"))
			auctioneer, err := model.GetAuctioneer()
			generatedUrl := service.URL()
			conductor := views.ConductAuction{auctioneer, generatedUrl, delay, bidAmount(0)}
			model.RegisterAuction(auctionData.AuctionId, conductor)
			// bidders, err := lib.FetchBidders()
			// service.MaxBidder(auctioneer, auctionData.AuctionId)
			render.JSON(w, err, conductor)
		}
	}
}
