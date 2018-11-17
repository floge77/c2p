package musiccloud_test

import (
	"cloud2podcast/musiccloud"
	"testing"
)

func TestPodcastInfo(t *testing.T) {
	tables := []struct {
		channelURL, expectedProvider string
	}{
		{"https://www.youtube.com/qdancedotnl", "youtube"},
		{"https://soundcloud.com/clnpodcast", "soundcloud"},
		{"https://asd.com/asdCast", "undefinedProvider"},
	}
	for _, table := range tables {
		p := musiccloud.NewPodcastinfo(
			"aChannel",
			table.channelURL,
			"anImageURL",
			"aPlaylisttoDownloadURL")

		if p.Provider != table.expectedProvider {
			t.Errorf("For %v Provider should be: %v, but was: %v", p.ChannelURL, table.expectedProvider, p.Provider)
		}
	}
}
