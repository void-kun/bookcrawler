package crawl

type MetruyencvSource struct {
	SourceType
}

func (s *MetruyencvSource) New(url string) {
	s.Name = "metruyencv"
	s.URL = url
}

func (s *MetruyencvSource) Search() {}
