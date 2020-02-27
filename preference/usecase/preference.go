package usecase

import "github.com/muathendirangu/lavida-api/preference"

type service struct {
	preferenceRepo preference.Repository
}

//NewService creates a new usecase/service
func NewService(repo preference.Repository) preference.Service {
	return &service{
		preferenceRepo: repo,
	}
}

func (s service) AddPreference(preference *preference.Preference) error {
	return s.preferenceRepo.Add(preference)
}
