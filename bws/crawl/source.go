package crawl

type Category struct {
	Name  string `json:"name"`
	Kinds string `json:"kinds"`
}

func (g *Category) New() {}

type SourceInterface interface {
	SearchBook()
	CrawlCategories()
}

type Source struct {
	Name       string     `json:"name"`
	URL        string     `json:"url"`
	Categories []Category `json:"categories"`
	Books      []Book     `json:"books"`
}
