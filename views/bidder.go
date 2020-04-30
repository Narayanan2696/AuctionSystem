package views

type BidderResponse struct {
	BidderDetails Bidder  `json:"bidder"`
	BidValue      float64 `json:"bid_value"`
}

type Bidder struct {
	Id   int    `json:"bidder_id"`
	Name string `json:"bidder_name"`
}
