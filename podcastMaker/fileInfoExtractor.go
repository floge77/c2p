package podcastMaker

import (
	"cloud2podcast/musiccloud"
	"os"
	"strings"
	"time"
)

type FileInfoExtractor struct {
}

func NewFileInfoExtractor() *FileInfoExtractor {
	return &FileInfoExtractor{}
}

func (f *FileInfoExtractor) GetPodcastItemsInformationForDir(dir string) (itemInfos []*musiccloud.PodcastItem, err error) {

	fileNames, err := f.readdir(dir)
	if err != nil {
		return nil, err
	}
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

func (*FileInfoExtractor) readdir(dirname string) (list []string, err error) {
	file, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	list, err = file.Readdirnames(0) // 0 to read all files and folders
	return
}
