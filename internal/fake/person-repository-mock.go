package fake

import (
	"poc-auto-read-ec/models"

	"github.com/stretchr/testify/mock"
)

type MockPersonRepository struct {
	mock.Mock
}

func (mock *MockPersonRepository) InsertPersonToDB(list []models.GormPerson) error {
	result := mock.Called(list)
	return result.Error(0)
}
