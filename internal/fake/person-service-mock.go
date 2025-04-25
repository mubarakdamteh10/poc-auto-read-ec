package fake

import (
	"poc-auto-read-ec/models"

	"github.com/stretchr/testify/mock"
)

type MockPersonService struct {
	mock.Mock
}

func (mock *MockPersonService) SavePersonsToDB(list []models.GormPerson) error {
	result := mock.Called(list)
	return result.Error(0)
}
