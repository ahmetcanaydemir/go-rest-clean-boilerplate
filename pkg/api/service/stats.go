package service

import (
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/api/repository"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/model"
)

type StatsService interface {
	GetStats() *model.Stats
}

type statsService struct {
	repository repository.StatsRepository
}

func NewStatsService(repositories ...repository.StatsRepository) StatsService {
	var repo repository.StatsRepository
	if len(repositories) == 0 {
		repo = repository.NewStatsRepository()
	} else {
		repo = repositories[0]
	}

	return &statsService{
		repository: repo,
	}
}

func (s statsService) GetStats() *model.Stats {
	stats := s.repository.GetLastStats()
	if len(stats) == 0 {
		return nil
	}
	stat := stats[0]

	return &stat
}
