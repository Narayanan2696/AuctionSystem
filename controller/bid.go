package controller

import (
	"auction_system/lib/render"
	"auction_system/model"
	"net/http"

	"github.com/gorilla/mux"
)

func WinningBid() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			request := mux.Vars(r)
			winningBid, err := model.GetWinningBid(request["auction_id"])
			render.JSON(w, err, winningBid)
		}
	}
}
