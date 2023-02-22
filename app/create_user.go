package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CreateUserRequest struct {
	Name string `json:"name"`
}

func (ctx *HandlerContext) CreateUser(w http.ResponseWriter, r *http.Request) {

	req := CreateUserRequest{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	newUserID, err := ctx.store.CreateUser(req.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%d", newUserID)
}
