package controller

import (
	"auction_system/lib"
	"auction_system/lib/render"
	"auction_system/model"
	"auction_system/views"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

func GetConductAuctioneer() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			request := mux.Vars(r)
			delay, _ := strconv.Atoi(os.Getenv("DELAY_MS"))
			bidders, _ := lib.FetchBidders()
			var response []views.ActiveBidderResponse
			bidderChannel := make(chan []views.ActiveBidderResponse)
			go getBidder(bidderChannel, bidders, delay, request["auction_id"])
			response = <-bidderChannel
			max := 0.0
			bidderId := 0
			for _, value := range response {
				if value.BidAmount > max {
					max = value.BidAmount
					bidderId = value.BidderId
				}
			}
			highestBidder := views.Bid{bidderId, max}
			model.InsertWinningBid(highestBidder, request["auction_id"])
			// auctioneer, err := model.GetAuctioneer()
			// generatedUrl := service.URL()
			// conductor := views.ConductAuction{auctioneer, generatedUrl, delay, bidAmount()}
			render.JSON(w, nil, highestBidder)
		}
	}
}

func getBidder(bidderChannel chan []views.ActiveBidderResponse, bidders []views.Bidder, delay int, auctionId string) {
	timeout := time.Tick(time.Duration(delay) * time.Millisecond)
	var activeBid []views.ActiveBidderResponse
	for {
		select {
		case <-timeout:
			bidderChannel <- activeBid
		default:
			liveAuction := model.GetLiveAuction(auctionId)
			bidder_id := getIndex(bidders)
			body := views.RegisterBidder{bidder_id, liveAuction.Url}
			lib.RegisterBidders(body, auctionId)
			activeBid = append(activeBid, views.ActiveBidderResponse{auctionId, bidder_id, "online", bidAmount(liveAuction.BidValue)})
		}
	}
}

func getIndex(bidders []views.Bidder) int {
	rand.Seed(time.Now().UnixNano())
	id := rand.Intn(len(bidders)-0) + 1
	return id
}

func bidAmount(bidVal float64) float64 {
	min_bid, _ := strconv.ParseFloat(os.Getenv("MIN_BID"), 64)
	max_bid, _ := strconv.ParseFloat(os.Getenv("MAX_BID"), 64)
	if bidVal > min_bid {
		min_bid = bidVal
	}
	rand.Seed(time.Now().UnixNano())
	bid := min_bid + rand.Float64()*(max_bid-min_bid)
	bid = (math.Round(bid*100) / 100)
	return bid
}
