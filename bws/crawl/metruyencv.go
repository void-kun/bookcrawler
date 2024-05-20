package crawl

type MetruyencvSource struct {
	Source
}

func (s *MetruyencvSource) New(url string) {
	s.Source.Name = "metruyencv"
	s.Source.URL = url
}

func (s *MetruyencvSource) Search() {}
