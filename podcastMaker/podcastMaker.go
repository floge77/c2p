package podcastMaker

import (

	"fmt"
	"github.com/floge77/cloud2podcast/musiccloud"

	"github.com/eduncan911/podcast"
)

type PodcastMaker struct {
}

func NewPodcastMaker() *PodcastMaker {
	return &PodcastMaker{}
}

func (*PodcastMaker) GetInitializedPodcast(podcastinfo *musiccloud.Podcastinfo) *podcast.Podcast {

	channel := podcastinfo.Channel
	provider := podcastinfo.Provider
	imageURL := podcastinfo.ChannelImageURL
	title := channel + "-" + provider + "-Podcast"

	p := podcast.New(
		title,
		podcastinfo.ChannelURL,
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
