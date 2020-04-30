package controller

import (
	"auction_system/lib/render"
	"auction_system/model"
	"auction_system/views"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterBidders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			request := mux.Vars(r)
			var register views.RegisterBidder
			json.NewDecoder(r.Body).Decode(&register)
			model.RegisterBidder(views.ActiveBidders{request["auction_id"], register.Id, "online"})
			render.JSON(w, nil, register)
		}
	}
}
