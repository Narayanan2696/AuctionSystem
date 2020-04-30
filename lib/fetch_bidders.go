package lib

import (
	"auction_system/lib/errors"
	"auction_system/views"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func FetchBidders() ([]views.Bidder, error) {
	host := os.Getenv("SOCKET")
	url := "http://" + host + "/bidders"
	req, err := http.Get(url)
	var bidders []views.Bidder
	if err != nil {
		log.Println(err.Error)
		return bidders, errors.New(errors.CustomError{404, "Not Found", "bidders not found"})
	}
	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &bidders)
	return bidders, err
}
