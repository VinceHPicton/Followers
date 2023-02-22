package app

import (
	"Followers/store"
	"net/http"

	"github.com/gorilla/mux"
)

type HandlerContext struct {
	store *store.Store
}

func NewHandlerContext(store *store.Store) *HandlerContext {
	if store == nil {
		panic("nil store!")
	}
	return &HandlerContext{store}
}

func Init(store *store.Store) *mux.Router {

	handlerCtx := NewHandlerContext(store)
	router := mux.NewRouter()

	// Create a user
	router.HandleFunc("/users", handlerCtx.CreateUser).Methods(http.MethodPost)

	// Follow a user
	router.HandleFunc("/users/{follower_id}/follow/{target_id}", handlerCtx.Follow).Methods(http.MethodPost)
	// Unfollow a user
	router.HandleFunc("/users/{follower_id}/unfollow/{target_id}", handlerCtx.Unfollow).Methods(http.MethodDelete)

	// Show all the people a user follows
	router.HandleFunc("/users/who/follow/{user_id}", handlerCtx.GetUsersWhoFollow).Methods(http.MethodGet)
	// Show all the people who follow the user
	router.HandleFunc("/followers/{user_id}", handlerCtx.GetFollowers).Methods(http.MethodGet)

	return router
}
