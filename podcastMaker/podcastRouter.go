package podcastMaker

import (
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

	config := configReader.GetConfig("/downloads/config.yaml")


	router := http.NewServeMux()
	c2p := NewCloud2podcast()

	router.Handle("/downloads/", http.StripPrefix("/downloads/", http.FileServer(http.Dir(config.DownloadDirectory))))
	for _, podcast := range config.PodcastsToServe {
		podcastInfo := musiccloud.NewPodcastinfo(
			podcast.Channel,
			podcast.ChannelURL,
			podcast.ChannelImageURL,
			podcast.PlaylistToDownloadURL)
		go router.HandleFunc(podcast.Channel ,func(w http.ResponseWriter, r *http.Request){
			c2p.ServePodcast(w, r, podcastInfo, config.DownloadDirectory)
		})
	}

	//router.HandleFunc("/podcasts", c2p.MakeAllPodcasts)
	log.Fatal(http.ListenAndServe(":8080", router))
}


