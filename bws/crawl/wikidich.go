package crawl

type WikidichSource struct {
	Source
}

func (s *WikidichSource) New(url string) {
	s.Source.Name = "wikidich"
	s.Source.URL = url
}

func (s *WikidichSource) Search() {}
