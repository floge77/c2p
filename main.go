package main

import "Cloud2Podcast/rss"

func main() {
	rss := rss.NewRssfeedBuilder()
	rss.PrintTemplate()
}
