package usecase

import "github.com/muathendirangu/lavida-api/trip"

type service struct {
	repo trip.Repository
}

//NewService creates the trip service/usecase
func NewService(tripRepo trip.Repository) trip.Service {
	return &service{
		repo: tripRepo,
	}
}

func (s *service) AddTrip(trip *trip.Trip) error {
	return s.repo.Store(trip)
}
