package podcastMaker

import (
	"fmt"
	"github.com/floge77/c2p/cloud2podcast/configReader"
	"github.com/floge77/c2p/cloud2podcast/musiccloud"
	"log"
	"net/http"
)

// reads config, creates routes to defined podcasts, passes config to podcastMaker
type PodcastRouter struct {
}

func NewPodcastRouter() *PodcastRouter {
	return &PodcastRouter{}
}

func (*PodcastRouter) ServePodcasts() {

	//config := configReader.GetConfig("/downloads/config.yaml")
	config := configReader.GetConfig("/Users/d068994/Development/private/c2p/cloud2podcast/config/config.yaml")


	router := http.NewServeMux()
	c2p := NewCloud2podcast()

	//router.Handle("/downloads/", http.StripPrefix("/downloads/", http.FileServer(http.Dir("/Users/d068994/Development/private/c2p/cloud2podcast/downloads/"))))
	router.Handle("/downloads/", http.StripPrefix("/downloads/", http.FileServer(http.Dir(config.DownloadDirectory))))
	for _, podcast := range config.PodcastsToServe {
		podcastInfo := musiccloud.NewPodcastinfo(
			podcast.Channel,
			podcast.ChannelURL,
			podcast.ChannelImageURL,
			podcast.PlaylistToDownloadURL)
		handlerFunc := c2p.ServePodcast(podcastInfo, config.DownloadDirectory)
		router.HandleFunc("/" +podcast.Channel , handlerFunc)
	}

	//router.HandleFunc("/podcasts", c2p.MakeAllPodcasts)
	fmt.Println("Router running at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}


