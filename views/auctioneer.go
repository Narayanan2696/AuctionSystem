package views

type ConductAuction struct {
	Conductor Auctioneer `json:"auctioneer"`
	Url       string     `json:"auction_url"`
	DelayMS   int        `json:"delay"`
	BidValue  float64    `json:"bid_value"`
}

type Auctioneer struct {
	Id   int    `json:"auctioneer_id"`
	Name string `json:"auctioneer_name"`
}
