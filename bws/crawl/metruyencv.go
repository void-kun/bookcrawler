package crawl

import "fmt"

// define crawling function for metruyenchu
type Metruyencv struct {
	Source
}

func NewMetruyencv(name string, url string) *Metruyencv {
	return &Metruyencv{
		Source: Source{
			Name: name,
			URL:  url,
		},
	}
}

func (m *Metruyencv) SearchBook() {
	fmt.Println("search books metruyencv")
}

func (m *Metruyencv) CrawlCategories() {
	fmt.Println("Craw categories metruyencv")
}
