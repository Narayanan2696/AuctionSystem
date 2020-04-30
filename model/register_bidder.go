package model

import "fmt"

func RegisterBidder() {
	rows, _ := connect.Query("SELECT * FROM auction_register")
	count := 0
	for rows.Next() {
		count += 1
	}
	fmt.Println(count + 1)
	// id := count + 1
	defer rows.Close()
	// insertQ, err := connect.Query("INSERT INTO AUCTION_REGISTER VALUES(?,?,?,?)", id, auctionId, location.Longitude)
}
