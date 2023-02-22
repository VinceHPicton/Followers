package main

import (
	"Followers/app"
	"Followers/store"
	"flag"
	"fmt"
	"net/http"
)

func main() {

	var httpPort int

	flag.IntVar(&httpPort, "http-port", 80, "HTTP server port")

	store := store.Start()
	router := app.Init(store)

	serverAddr := fmt.Sprintf(":%v", httpPort)
	http.ListenAndServe(serverAddr, router)
}
