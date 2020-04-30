package views

type ActiveBidders struct {
	AuctionId string `json:"auction_id"`
	Id        int    `json:"bidder_id"`
	Name      string `json:"bidder_name"`
	Url       string `json:"auction_url"`
	Status    string `json:"status"`
}
