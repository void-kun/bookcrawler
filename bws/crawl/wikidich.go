package crawl

import "fmt"

// define crawling function for wikidich
type Wikidich struct {
	Source
}

func NewWikidich(name string, url string) *Wikidich {
	return &Wikidich{
		Source: Source{
			Name: name,
			URL:  url,
		},
	}
}

func (w *Wikidich) SearchBook() {
	fmt.Println("search books")
}

func (w *Wikidich) CrawlCategories() {
	fmt.Println("Craw categories")
}
