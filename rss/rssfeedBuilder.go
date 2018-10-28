package rss

import (
	"fmt"
	"os"
)

type RssfeedBuilder struct {
}

func NewRssfeedBuilder() *RssfeedBuilder {
	RssfeedBuilder := &RssfeedBuilder{}
	return RssfeedBuilder
}

func (*RssfeedBuilder) PrintTemplate() {
	xmlFile, err := os.Open("rss/podcastTemplate.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()
}
