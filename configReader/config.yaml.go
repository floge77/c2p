package configReader

import (
	"time"
)

type PodcastConfig struct {
	DownloadDirectory     string                    `yaml:"downloadDirectory"`
	MinSetLengthInSeconds int                       `yaml:"minSetLengthInSeconds"`
	PodcastsToServe       []*Podcastinfo `yaml:"podcasts"`
}

type PodcastItem struct {
	Title, Channel, FileName string
	FileSize                 int64
	ReleaseDate              *time.Time
}

type Podcastinfo struct {
	Channel               string `yaml:"channelName"`
	ChannelURL            string `yaml:"channelURL"`
	ChannelImageURL       string `yaml:"channelImageURL"`
	PlaylistToDownloadURL string `yaml:"playlistToDownloadURL"`
	Provider              string
	DownloadDirectory     string
	Items                 []*PodcastItem
}