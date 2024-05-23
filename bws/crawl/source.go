package crawl

type Gender struct {
	Name  string `json:"name"`
	Kinds string `json:"kinds"`
}

func (g *Gender) New() {}

type SourceType struct {
	Name    string   `json:"name"`
	URL     string   `json:"url"`
	Genders []Gender `json:"genders"`
	Books   []Book   `json:"books"`
}

// Generics Source type
type SourceParse interface {
	*WikidichSource | *MetruyencvSource
	Search()
}

type Source[T SourceParse] struct {
	Body T
}

func New[T SourceParse](object T) *Source[T] {
	return &Source[T]{
		Body: object,
	}
}

func (s *Source[T]) Search() {
	s.Body.Search()
}
