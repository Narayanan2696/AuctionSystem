package controller

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func ListOnlineBidders() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := mux.Vars(r)
		fmt.Println(request["auction_id"])
	}
}
