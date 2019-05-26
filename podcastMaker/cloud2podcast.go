package podcastMaker

import (
	"github.com/floge77/cloud2podcast/musiccloud"
	"log"
	"net/http"
	"os"
)

type Cloud2podcast struct {
	podcastMaker      *PodcastMaker
	fileInfoExtractor *FileInfoExtractor
}

func NewCloud2podcast() *Cloud2podcast {
	c2p := &Cloud2podcast{}
	c2p.podcastMaker = NewPodcastMaker()
	c2p.fileInfoExtractor = NewFileInfoExtractor()
	return c2p
}

func (c *Cloud2podcast) ServePodcast(podcastInfo *musiccloud.Podcastinfo, generalDownloadDirectory string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		completeDownloadDirectory := generalDownloadDirectory + "/" + podcastInfo.Provider + "/" + podcastInfo.Channel
		var err error

		podcastInfo.Items, err = c.fileInfoExtractor.GetPodcastItemsInformationForDir(completeDownloadDirectory)
		if err != nil {
			log.Printf("Could not serve Podcast: %v Error: %v", podcastInfo.Channel, err)
		} else {
			podcast := c.podcastMaker.GetInitializedPodcast(podcastInfo)

			hostIP  := os.Getenv("HOST_IP")
			
			for _, item := range podcastInfo.Items {
				c.podcastMaker.AppendPodcastItem(podcast, item, "http://" + hostIP + ":8080" + completeDownloadDirectory + "/")
				//c.podcastMaker.AppendPodcastItem(podcast, item, "http://192.168.178.36:8080"+completeDownloadDirectory+"/")
			}

			w.Header().Set("Content-Type", "application/xml")

			if err := podcast.Encode(w); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
		}
	}
}
