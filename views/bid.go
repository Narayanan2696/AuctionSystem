package views

type Bid struct {
	BidderId      int     `json:"bidder_id"`
	BiddingAmount float64 `json:"bidding_amount"`
}

type WinningBid struct {
	Id            int
	BidderId      int
	AuctionId     string
	BiddingAmount float64
}
