package podcastMaker

import (
	"cloud2podcast/configReader"
	"cloud2podcast/musiccloud"
	"log"
	"net/http"
)

type Cloud2podcast struct {
	podcastMaker      *PodcastMaker
	fileInfoExtractor *FileInfoExtractor
}

//func Handle() http.Handler {
//	return http.HandlerFunc(MakeAllPodcasts)
//}

func GetConfig(yamlPath string) configReader.PodcastConfig {
	yamlReader := configReader.NewYamlreader()
	config := configReader.PodcastConfig{}
	config = yamlReader.ReadYamlfile(yamlPath, config)
	return config
}

func (c *Cloud2podcast) MakeAllPodcasts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	c.podcastMaker = NewPodcastMaker()
	c.fileInfoExtractor = NewFileInfoExtractor()

	config := GetConfig("config/config.yaml")

	for _, podcast := range config.PodcastsToServe {
		podcastInfo := musiccloud.NewPodcastinfo(
			podcast.Channel,
			podcast.ChannelURL,
			podcast.ChannelImageURL,
			podcast.PlaylistToDownloadURL)
		go c.servePodcast(w, r, podcastInfo, config.DownloadDirectory)
	}
}

func (c *Cloud2podcast) servePodcast(w http.ResponseWriter, r *http.Request, podcastInfo *musiccloud.Podcastinfo, generalDownloadDirectory string) {
	completeDownloadDirectory := generalDownloadDirectory + "/" + podcastInfo.Provider + "/" + podcastInfo.Channel
	var err error

	podcastInfo.Items, err = c.fileInfoExtractor.GetPodcastItemsInformationForDir(completeDownloadDirectory)
	if err != nil {
		log.Printf("Could not serve Podcast: %v Error: %v", podcastInfo.Channel, err)
	} else {
		podcast := c.podcastMaker.GetInitializedPodcast(podcastInfo)

		for _, item := range podcastInfo.Items {
			c.podcastMaker.AppendPodcastItem(podcast, item, "http://192.168.178.30:8080/"+completeDownloadDirectory+"/")
		}

		w.Header().Set("Content-Type", "application/xml")

		if err := podcast.Encode(w); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}

}
