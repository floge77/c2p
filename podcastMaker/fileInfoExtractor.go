package podcastMaker

import (
	"Cloud2Podcast/musiccloud"
	"log"
	"os"
	"strings"
	"time"
)

type FileInfoExtractor struct {
}

func NewFileInfoExtractor() *FileInfoExtractor {
	return &FileInfoExtractor{}
}

func (f *FileInfoExtractor) GetPodcastItemsInformationForDir(dir string) (itemInfos []*musiccloud.PodcastItem) {

	fileNames := f.readdir(dir)
	for _, name := range fileNames {
		// create Info structs for podcastsItems
		item := f.getPodcastItemInfosFromFileName(dir, name)
		itemInfos = append(itemInfos, item)
	}
	return
}

func (f *FileInfoExtractor) getPodcastItemInfosFromFileName(dir string, filename string) (item *musiccloud.PodcastItem) {

	s := strings.Replace(filename, ".mp3", "", -1)
	fields := strings.Split(s, "__")
	item = &musiccloud.PodcastItem{}
	item.Title = fields[1]
	item.Channel = fields[2]
	item.ReleaseDate = f.getReleaseDateFromString(fields[0])
	fileSize, _ := f.extractFileSize(dir, filename)
	item.FileSize = fileSize
	item.FileName = filename
	return
}

func (*FileInfoExtractor) extractFileSize(dir string, filename string) (fileSize int64, err error) {
	file, err := os.Stat(dir + "/" + filename)
	if err != nil {
		return 0, err
	}
	return file.Size(), nil
}

func (*FileInfoExtractor) getReleaseDateFromString(date string) *time.Time {
	t, _ := time.Parse("20060102", date)
	return &t
}

func (*FileInfoExtractor) readdir(dirname string) []string {
	file, err := os.Open(dirname)
	if err != nil {
		log.Fatalf("failed opening directory: %s", err)
	}
	defer file.Close()
	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	return list
}
