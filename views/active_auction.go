package views

type ActiveAuction struct {
	AuctionId    string `json:"auction_id"`
	AuctioneerId int    `json:"auctioneer_id"`
	Url          string `json:"auction_url"`
}
