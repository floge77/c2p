package main

import "github.com/floge77/cloud2podcast/podcastMaker"

func main() {
	podcastRouter := podcastMaker.NewPodcastRouter()
	podcastRouter.ServePodcasts()
}
