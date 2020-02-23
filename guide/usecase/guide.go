package usecase

import "github.com/muathendirangu/lavida-api/guide"

type service struct {
	repo guide.Repository
}

//NewService creates the guide service/usecase
func NewService(guideRepo guide.Repository) guide.Service {
	return &service{
		repo: guideRepo,
	}
}

func (s *service) AddGuide(guide *guide.Guide) error {
	return s.repo.Store(guide)
}

func (s *service) GetGuides() []*guide.Guide {
	return s.repo.Get()
}
