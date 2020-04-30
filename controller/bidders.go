package controller

import (
	"auction_system/lib/render"
	"auction_system/model"
	"fmt"
	"net/http"
)

func GetBidders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			bidders, err := model.GetAllBidders()
			fmt.Println("reached bidders:")
			render.JSON(w, err, bidders)
		}
	}
}
