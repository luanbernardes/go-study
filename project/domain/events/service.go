package events

import "first-app/models"

type (
	Servicer interface {
		GetAllEvents() []models.Event
	}
	ServiceParams struct {
		Repo Repository
	}
	Service struct {
		repo Repository
	}
)

func NewService(params ServiceParams) Servicer {
	return &Service{
		repo: params.Repo,
	}
}

func (s *Service) GetAllEvents() []models.Event {
	return s.repo.GetAllEvents()
}
