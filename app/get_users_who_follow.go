package app

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (ctx *HandlerContext) GetUsersWhoFollow(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userIDStr := vars["user_id"]

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	usernamesFollowedByUser := ctx.store.GetUsersWhoFollow(userID)

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%v", usernamesFollowedByUser)
}
