package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ctx *HandlerContext) Unfollow(w http.ResponseWriter, r *http.Request) {

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

	idDeleted, err := ctx.store.DeleteFollow(followerID, targetID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%v", idDeleted)
}
