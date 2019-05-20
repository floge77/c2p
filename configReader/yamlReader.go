package configReader

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type YamlReader struct {
}

func NewYamlreader() *YamlReader {
	return &YamlReader{}
}

func GetConfig(yamlPath string) PodcastConfig {
	yamlReader := NewYamlreader()
	config := PodcastConfig{}
	config = yamlReader.ReadYamlfile(yamlPath, config)
	return config
}

func (*YamlReader) ReadYamlfile(filePath string, config PodcastConfig) PodcastConfig {

	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Could not open %v Error: %v", filePath, err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", config)
	return config
}
