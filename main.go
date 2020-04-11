package main

import (
	"d2g-server/presenter"
	"log"
	"net/http"
)

func main() {
	server := presenter.SetupSocket()
	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))
	log.Println("Serving at localhost:6060...")
	log.Fatal(http.ListenAndServe(":6060", nil))
}
