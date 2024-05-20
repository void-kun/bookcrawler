package crawl

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

// init a struct for crawler
// load the config for sources
// search book by (keyword, source)

type Metadata struct {
	Url string
}

type Crawl struct {
	Sources []Source
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
		var wikidichSource SourceParse
		wikidichSource = WikidichSource{}
		wikidichSource.New(metadata["wikidich"].Url)
		c.Sources = append(c.Sources, wikidichSource)
	}

	if metadata["metruyencv"] != (Metadata{}) {
		metruyencvSource := MetruyencvSource{}
		metruyencvSource.New(metadata["wikidich"].Url)
		c.Sources = append(c.Sources, metruyencvSource)
	}

	return nil
}

func (c *Crawl) Search() {}

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
