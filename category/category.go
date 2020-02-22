package category

import (
	"github.com/muathendirangu/lavida-api/domains"
)

type service struct {
	repo domains.Repository
}

//NewService creates the category service/usecase
func NewService(userRepo domains.Repository) Service {
	return &service{
		repo: userRepo,
	}
}

func (s *service) AddCategory(category *domains.Category) error {
	return s.repo.Store(category)
}

func (s *service) GetCategories() []*domains.Category {
	return s.repo.Get()
}
