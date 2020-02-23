package category

import "github.com/muathendirangu/lavida-api/category"

type service struct {
	repo category.Repository
}

//NewService creates the category service/usecase
func NewService(userRepo category.Repository) category.Service {
	return &service{
		repo: userRepo,
	}
}

func (s *service) AddCategory(category *category.Category) error {
	return s.repo.Store(category)
}

func (s *service) GetCategories() []*category.Category {
	return s.repo.Get()
}

func (s *service) GetCategory(ID int) *category.Category {
	return s.repo.GetByID(ID)
}
