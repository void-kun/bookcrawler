package crawl

type WikidichSource struct {
	SourceType
}

func (s *WikidichSource) New(url string) {
	s.Name = "wikidich"
	s.URL = url
}

func (s *WikidichSource) Search() {}
