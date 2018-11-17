package musiccloud

import (
	"errors"
	"net/url"
	"strings"
	"time"
)

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

func NewPodcastinfo(channel, channelURL, channelImageURL, playlistToDownloadURL string) *Podcastinfo {
	p := &Podcastinfo{}
	p.Channel = channel
	p.ChannelURL = channelURL
	p.ChannelImageURL = channelImageURL
	p.PlaylistToDownloadURL = playlistToDownloadURL

	provider, err := p.getProviderFromChannelURL(channelURL)
	if err != nil {
		p.Provider = "undefinedProvider"
	} else {
		p.Provider = provider
	}
	p.DownloadDirectory = p.Provider + "/" + p.Channel + "/"
	return p
}

func (p *Podcastinfo) getProviderFromChannelURL(channelURL string) (provider string, err error) {

	u, err := url.Parse(channelURL)
	provider = strings.Replace(u.Hostname(), "www.", "", -1)
	provider = strings.Split(provider, ".")[0]

	err = checkProviderIsAllowed(provider)
	return
}

func checkProviderIsAllowed(provider string) (err error) {
	var allowedProviders = []string{"soundcloud", "mixcloud", "youtube"}
	providerNotFound := true
	for _, ap := range allowedProviders {
		if strings.Contains(provider, ap) {
			providerNotFound = false
		}
	}
	if providerNotFound {
		return errors.New("Provider not in: " + strings.Join(allowedProviders, ","))
	} else {
		return nil
	}
}
