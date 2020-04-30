package lib

import (
	"auction_system/views"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func RegisterBidders(body views.RegisterBidder, auctionId string) {
	host := os.Getenv("SOCKET")
	url := "http://" + host + "/register_bidders/" + auctionId
	requestByte, _ := json.Marshal(body)
	requestReader := bytes.NewReader(requestByte)
	res, err := http.Post(url, "application/json", requestReader)
	if res.StatusCode == 200 {
		fmt.Println(body.Url)
	} else {
		log.Println(res.Status)
		log.Println(err)
	}
}
