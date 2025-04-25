package fake

import (
	"poc-auto-read-ec/models"

	"github.com/stretchr/testify/mock"
)

type MockJSONToSoapService struct {
	mock.Mock
}

func (mock *MockJSONToSoapService) CreateSoapData() (*models.SoapData, error) {
	return &models.SoapData{
		
	}, nil
}
