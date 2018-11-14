package main

import (
	"Cloud2Podcast/podcastMaker"
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.Handle("/downloads/", http.StripPrefix("/downloads/", http.FileServer(http.Dir("downloads"))))
	router.Handle("/podcasts", podcastMaker.Handle())
	log.Fatal(http.ListenAndServe(":8080", router))

}
