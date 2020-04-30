package views

type LiveAuction struct {
	Id           int
	AuctionId    string
	AuctioneerId int
	Url          string
	BidValue     float64
}
type ActiveAuction struct {
	AuctionId    string `json:"auction_id"`
	AuctioneerId int    `json:"auctioneer_id"`
	Url          string `json:"auction_url"`
}
