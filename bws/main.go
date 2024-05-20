package main

import (
	"bws/crawl"
)

func main() {
	metadataPath := "./sources.json"
	myCrawl := crawl.Crawl{}

	myCrawl.LoadSources(metadataPath)
}
