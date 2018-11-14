package musiccloud

import "time"

// will be implemented by mixcloud soundcloud youtube whatever
type CloudMusicProvider interface {
	DownloadTracks()
	GetChannel() string
	GetProvider() string
	GetChannelURL() string
	GetChannelImageURL() string
	GetPlaylistToDownloadURL() string
}

type PodcastItem struct {
	Title, Channel, FileName string
	FileSize                 int64
	ReleaseDate              *time.Time
}
