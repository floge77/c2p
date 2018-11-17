package main

import (
	"cloud2podcast/podcastMaker"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	c2p := podcastMaker.Cloud2podcast{}
	router.Handle("/downloads/", http.StripPrefix("/downloads/", http.FileServer(http.Dir("/downloads"))))
	router.HandleFunc("/podcasts", c2p.MakeAllPodcasts)
	log.Fatal(http.ListenAndServe(":8080", router))
}
