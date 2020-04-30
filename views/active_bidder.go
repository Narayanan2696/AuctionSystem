package views

type ActiveBidderResponse struct {
	AuctionId string  `json:"auction_id"`
	BidderId  int     `json:"bidder_id"`
	Status    string  `json:"status"`
	BidAmount float64 `json:"bid_amount"`
}

type RegisterBidder struct {
	Id  int    `json:"bidder_id"`
	Url string `json:"auction_url"`
}

type ActiveBidders struct {
	AuctionId string `json:"auction_id"`
	BidderId  int    `json:"bidder_id"`
	// Url       string `json:"auction_url"`
	Status string `json:"status"`
}
