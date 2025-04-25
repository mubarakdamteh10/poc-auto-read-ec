package person

import (
	"errors"
	"poc-auto-read-ec/internal/fake"
	"poc-auto-read-ec/models"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPersonService(t *testing.T) {
	service := NewPersonService()

	v := reflect.Indirect(reflect.ValueOf(service))
	for index := 0; index < v.NumField(); index++ {
		assert.False(t, v.Field(index).IsZero(), "Field %s is zero value", v.Type().Field(index).Name)
	}

	assert.NotNil(t, service, "Service should not be nil")
}

func TestSavePersonsToDB(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		// Arrange

		mockPerson := []models.GormPerson{
			{
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "yourMomIsGay@male.com",
				PhoneNumber: "0123456789",
				DateOfBirth: "1990-01-01",
				Address:     "123 Main St",
			},
			{
				FirstName:   "Jane",
				LastName:    "MaiDee",
				Email:       "yourMomIsGay@male.com",
				PhoneNumber: "0123456789",
				DateOfBirth: "1990-01-01",
				Address:     "123 Main St",
			},
		}

		mockPersonService := &fake.MockPersonRepository{}
		mockPersonService.On("InsertPersonToDB", mockPerson).Return(nil)

		service := personService{
			repository: mockPersonService,
		}
		// Act

		err := service.SavePersonsToDB(mockPerson)

		// Assert
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		// Arrange

		mockPerson := []models.GormPerson{
			{
				FirstName:   "John",
				LastName:    "Doe",
				Email:       "yourMomIsGay@male.com",
				PhoneNumber: "0123456789",
				DateOfBirth: "1990-01-01",
				Address:     "123 Main St",
			},
			{
				FirstName:   "Jane",
				LastName:    "MaiDee",
				Email:       "yourMomIsGay@male.com",
				PhoneNumber: "0123456789",
				DateOfBirth: "1990-01-01",
				Address:     "123 Main St",
			},
		}

		mockPersonService := &fake.MockPersonRepository{}
		mockPersonService.On("InsertPersonToDB", mockPerson).Return(errors.New("waiting for implement"))

		service := personService{
			repository: mockPersonService,
		}
		// Act

		err := service.SavePersonsToDB(mockPerson)

		// Assert
		assert.ErrorContains(t, err, "waiting for implement")
	})
}
