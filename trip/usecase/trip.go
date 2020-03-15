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

func (s *service) SearchTrip(Traveler, Month, Activity, Category string) []*trip.Trip {
	return s.repo.Search(Traveler, Month, Activity, Category)
}

func (s *service) ViewTrip(ID, Guide int64) *trip.Trip {
	return s.repo.View(ID, Guide)
}

func (s *service) GetTrips() []*trip.Trip {
	return s.repo.Get()
}
