package musiccloud

import (
	"fmt"
	"log"
	"os/exec"
)

const youtubeDownloader = "youtube-dl"
const youtubeDownloaderConfigPath = "config/youtube-dl.conf"

type Youtube struct {
	Channel, ChannelURL, ChannelImageURL, PlaylistToDownloadURL, provider string
	Items                                                                 []*PodcastItem
}

func NewYoutube() *Youtube {
	youtube := &Youtube{}
	youtube.provider = "Youtube"
	return youtube
}

// This method is the public API
// TODO:
// receive output of youtube-dl on console
// modify youtube-dl.conf to download to "downloads/provider/channel"
func (y *Youtube) DownloadTracks() {
	cmdString := youtubeDownloader + " --config-location " + youtubeDownloaderConfigPath + " " + y.ChannelURL
	cmd := exec.Command("sh", "-c", cmdString)
	//cmd = exec.Command(youtubeDownloader, "--config-location", youtubeDownloaderConfigPath, y.channelURL)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}

func (y *Youtube) GetChannel() string {
	return y.Channel
}
func (y *Youtube) GetProvider() string {
	return y.provider
}
func (y *Youtube) GetChannelURL() string {
	return y.ChannelURL
}
func (y *Youtube) GetChannelImageURL() string {
	return y.ChannelImageURL
}
func (y *Youtube) GetPlaylistToDownloadURL() string {
	return y.PlaylistToDownloadURL
}
