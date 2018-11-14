package podcastMaker

import (
	"Cloud2Podcast/musiccloud"
	"net/http"
)

type Cloud2Podcast struct {
}

//func NewCloud2Podcast() *Cloud2Podcast {
//	return &Cloud2Podcast{}
//}

func Handle() http.Handler {
	return http.HandlerFunc(MakePodcasts)
}

func MakePodcasts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	youtube := musiccloud.NewYoutube()
	podcastMaker := NewPodcastMaker()
	fileInfoExtractor := NewFileInfoExtractor()

	youtube.Channel = "Q-Dance"
	youtube.ChannelURL = "https://www.youtube.com/qdancedotnl"
	youtube.ChannelImageURL = "https://yt3.ggpht.com/a-/AN66SAzyW12uAQRayPY4MS_Fo_Wlj6PFjyNfx3X7CQ=s288-mo-c-c0xffffffff-rj-k-no"
	youtube.PlaylistToDownloadURL = "https://www.youtube.com/watch\\?v\\=ZMPlY-FtJIM\\&list\\=UUAEwCfBRlB3jIY9whEfSP5Q"
	youtube.Items = fileInfoExtractor.GetPodcastItemsInformationForDir("downloads")

	youtubePodcast := podcastMaker.GetInitializedPodcast(youtube)
	for _, item := range youtube.Items {
		podcastMaker.AppendPodcastItem(youtubePodcast, item, "http://192.168.178.30:8080/downloads/")
	}

	w.Header().Set("Content-Type", "application/xml")

	if err := youtubePodcast.Encode(w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
