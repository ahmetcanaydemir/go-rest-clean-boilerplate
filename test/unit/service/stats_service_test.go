package service_test

import (
	"testing"

	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/api/service"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/pkg/model"
	"github.com/ahmetcanaydemir/go-rest-clean-boilerplate/test/unit/repository/mock"
	"github.com/stretchr/testify/assert"
)

func Test_GetStats(t *testing.T) {
	// Arrange
	stats := []model.Stats{
		{
			Count: 1,
		},
		{
			Count: 2,
		},
	}
	statsRepositoryMock := new(mock.StatsRepositoryMock)
	statsRepositoryMock.On("GetLastStats").Return(stats)
	statsService := service.NewStatsService(statsRepositoryMock)

	// Act
	stat := statsService.GetStats()

	// Assert
	statsRepositoryMock.AssertExpectations(t)
	assert.Equal(t, stats[0], *stat)
}

func Test_GetStats_WithEmptyStats(t *testing.T) {
	// Arrange
	var expectedStat *model.Stats
	var stats []model.Stats
	statsRepositoryMock := new(mock.StatsRepositoryMock)
	statsRepositoryMock.On("GetLastStats").Return(stats)
	statsService := service.NewStatsService(statsRepositoryMock)

	// Act
	stat := statsService.GetStats()

	// Assert
	statsRepositoryMock.AssertExpectations(t)
	assert.Equal(t, expectedStat, stat)
}
