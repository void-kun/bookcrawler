package main

import (
	"bws/crawl"
)

func main() {
	crawler := crawl.Crawl{}

	crawler.LoadSources("./sources.json")
	crawler.Search()
}
