package lib

import (
	"auction_system/views"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func FetchAuctioneer(auction_id string) (views.ConductAuction, error) {
	host := os.Getenv("SOCKET")
	url := "http://" + host + "/auctioneer"
	req, err := http.Get(url)
	var conductor views.ConductAuction
	if err != nil {
		log.Println(err.Error)
		return views.ConductAuction{}, err
	}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &conductor)

	return conductor, err
}
