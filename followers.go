package main

import (
	"Followers/user_service"
	"flag"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	var httpPort int

	flag.IntVar(&httpPort, "http-port", 80, "HTTP server port")

	router := mux.NewRouter()

	router.HandleFunc("/users", user_service.Create).Methods(http.MethodPost)
	router.HandleFunc("/users/{follower_id}/follow/{following_id}", user_service.NewFollow).Methods(http.MethodPost)

	serverAddr := fmt.Sprintf(":%v", httpPort)
	http.ListenAndServe(serverAddr, router)
}
