package main

import "github.com/floge77/c2p/cloud2podcast/podcastMaker"

func main() {
	podcastRouter := podcastMaker.NewPodcastRouter()
	podcastRouter.ServePodcasts()
}
