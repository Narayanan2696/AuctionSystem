package controller

import (
	"auction_system/lib/render"
	"auction_system/model"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func WinningBid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			request := mux.Vars(r)
			list, err := model.GetWinningBid(request["auction_id"])
			fmt.Println(list)
			render.JSON(w, err, list)
		}
	}
}
