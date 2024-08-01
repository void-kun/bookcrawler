package crawl

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

type Metadata struct {
	Url string
}

type Crawl struct {
	Sources []SourceInterface
}

func (c *Crawl) LoadSources(dataPath string) error {
	// load metadata from file
	metadataByte, err := parseJson(dataPath)
	if err != nil {
		return err
	}

	var metadata map[string]Metadata
	err = json.Unmarshal(metadataByte, &metadata)
	if err != nil {
		log.Fatal("Source config is error: ", err)
		return err
	}

	if metadata["wikidich"] != (Metadata{}) {
		wikidich := NewWikidich("wikidich", metadata["wikidich"].Url)
		c.Sources = append(c.Sources, wikidich)
	}

	if metadata["metruyencv"] != (Metadata{}) {
		metruyencv := NewMetruyencv("metruyencv", metadata["metruyencv"].Url)
		c.Sources = append(c.Sources, metruyencv)
	}

	return nil
}

func (c *Crawl) Search() {
	for _, item := range c.Sources {
		item.SearchBook()
	}
}

func (c *Crawl) Save() {}

func parseJson(jsonPath string) ([]byte, error) {
	jsonFile, err := os.Open(jsonPath)
	if err != nil {
		log.Fatal("Cannot open file: ", err)
		return nil, err
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal("Cannot read file: ", err)
		return nil, err
	}
	return byteValue, nil
}
