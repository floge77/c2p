package configReader

import "cloud2podcast/musiccloud"

type PodcastConfig struct {
	DownloadDirectory     string                    `yaml:"downloadDirectory"`
	MinSetLengthInSeconds int                       `yaml:"minSetLengthInSeconds"`
	PodcastsToServe       []*musiccloud.Podcastinfo `yaml:"podcasts"`
}
