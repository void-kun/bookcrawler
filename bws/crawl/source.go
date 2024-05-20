package crawl

type Gender struct {
	Name  string `json:"name"`
	Kinds string `json:"kinds"`
}

func (g *Gender) New() {}

type Source struct {
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Genders []Gender `json:"genders"`
	Books   []Book   `json:"books"`
}

type SourceParse interface {
	New()
	Search()
}
