package mock

import (
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/api/repository"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/model"
	"github.com/stretchr/testify/mock"
)

type StatsRepositoryMock struct {
	mock.Mock
}

var _ repository.StatsRepository = (*StatsRepositoryMock)(nil)

func (m *StatsRepositoryMock) GetLastStats() []model.Stats {
	args := m.Called()
	return args.Get(0).([]model.Stats)
}
