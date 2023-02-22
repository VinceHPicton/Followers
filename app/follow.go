package app

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ctx *HandlerContext) Follow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	followerIDStr := vars["follower_id"]
	targetIDStr := vars["target_id"]

	followerID, err := strconv.Atoi(followerIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	targetID, err := strconv.Atoi(targetIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = ctx.store.CreateFollow(followerID, targetID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
