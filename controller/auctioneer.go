package controller

import (
	"auction_system/lib/render"
	"auction_system/model"
	"auction_system/service"
	"auction_system/views"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func GetConductAuctioneer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			delay, _ := strconv.Atoi(os.Getenv("DELAY_MS"))
			auctioneer, err := model.GetAuctioneer()
			generatedUrl := service.URL()
			conductor := views.ConductAuction{auctioneer, generatedUrl, delay, bidAmount()}
			render.JSON(w, err, conductor)
		}
	}
}

func bidAmount() float64 {
	min_bid, _ := strconv.ParseFloat(os.Getenv("MIN_BID"), 64)
	max_bid, _ := strconv.ParseFloat(os.Getenv("MAX_BID"), 64)
	rand.Seed(time.Now().UnixNano())
	bid := min_bid + rand.Float64()*(max_bid-min_bid)
	bid = (math.Round(bid*100) / 100)
	return bid
}
