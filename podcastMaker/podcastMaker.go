package podcastMaker

import (
	"Cloud2Podcast/musiccloud"
	"fmt"

	"github.com/eduncan911/podcast"
)

type PodcastMaker struct {
}

func NewPodcastMaker() *PodcastMaker {
	return &PodcastMaker{}
}

func (*PodcastMaker) GetInitializedPodcast(cloudMusicProvider musiccloud.CloudMusicProvider) *podcast.Podcast {

	channel := cloudMusicProvider.GetChannel()
	provider := cloudMusicProvider.GetProvider()
	imageURL := cloudMusicProvider.GetChannelImageURL()
	title := channel + "-" + provider + "-Podcast"

	p := podcast.New(
		title,
		cloudMusicProvider.GetChannelURL(),
		"",
		nil, nil,
	)

	p.ISubtitle = title
	p.AddSummary(channel + " Podcast from " + channel + " " + provider + " channel")
	p.AddImage(imageURL)
	p.AddAuthor(channel, channel+"@email.com")

	return &p
}

func (*PodcastMaker) AppendPodcastItem(podcastToAppend *podcast.Podcast, itemToAdd *musiccloud.PodcastItem, fileServerURL string) {

	title := itemToAdd.Title
	filename := itemToAdd.FileName
	channel := itemToAdd.Channel
	releaseDate := itemToAdd.ReleaseDate
	fileSize := itemToAdd.FileSize

	item := podcast.Item{
		Title:       title,
		Description: channel,
		ISubtitle:   "",
		PubDate:     releaseDate,
	}
	item.AddSummary(title)
	item.AddEnclosure(fileServerURL+filename, podcast.MP3, fileSize)

	_, err := podcastToAppend.AddItem(item)
	if err != nil {
		fmt.Println(item.Title, ": error", err.Error())
	}
}
